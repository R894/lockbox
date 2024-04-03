package ssh

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"

	"github.com/R894/lockbox/internal/tunnel"

	"github.com/gliderlabs/ssh"
)

type LockBox struct {
	server *ssh.Server
	tm     *tunnel.TunnelManager
}

func NewServer(addr string, tm *tunnel.TunnelManager) *LockBox {
	return &LockBox{server: &ssh.Server{
		Addr: addr,
	},
		tm: tm,
	}
}

func (l *LockBox) ListenAndServe() error {
	fmt.Println("SSH Server listening on", l.server.Addr)
	ssh.Handle(func(s ssh.Session) {
		handleRequest(s, l.tm)
	})
	return l.server.ListenAndServe()
}

func (l *LockBox) Close() error {
	return l.server.Close()
}

func handleRequest(s ssh.Session, tm *tunnel.TunnelManager) {
	id := rand.Intn(math.MaxInt)
	s.Write(([]byte)(fmt.Sprintf("LockBox link: http://localhost:3000?id=%d\n", id)))
	s.Write(([]byte)("Session is in progress... Waiting for user to connect and download\n"))

	currentTunnel := tm.AddTunnel(id)
	donech := make(chan struct{})
	currentTunnel <- tunnel.Tunnel{
		Filename: s.RawCommand(),
		Donech:   donech,
	}

	tunnelChanAfterLinkClicked, _ := tm.GetTunnel(id)
	tunnelAfterLinkClicked := <-tunnelChanAfterLinkClicked
	defer close(tunnelAfterLinkClicked.Donech)
	err := sendFileToTunnel(s, &tunnelAfterLinkClicked)

	if err != nil {
		log.Fatal(err)
		s.Write(([]byte)("Something went wrong!\n"))
		return
	}
	s.Write(([]byte)("File sent successfully, thanks for using LockBox!\n"))
}

func sendFileToTunnel(session io.Reader, tunnel *tunnel.Tunnel) error {
	fmt.Println("Inside sendFileToTunnel")
	_, err := io.Copy(tunnel.W, session)
	if err != nil {
		return err
	}
	fmt.Println("File sent successfully")
	return nil
}
