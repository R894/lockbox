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
}

func NewServer(addr string, tm *tunnel.TunnelManager) *LockBox {
	ssh.Handle(func(s ssh.Session) {
		handleRequest(s, nil)
	})
	return &LockBox{server: &ssh.Server{
		Addr: addr,
		Handler: func(s ssh.Session) {
			handleRequest(s, tm)
		},
	}}
}

func (l *LockBox) ListenAndServe() error {
	return l.server.ListenAndServe()
}

func (l *LockBox) Close() error {
	return l.server.Close()
}

func handleRequest(s ssh.Session, tm *tunnel.TunnelManager) {
	id := rand.Intn(math.MaxInt)
	tunnelChan := tm.AddTunnel(id)
	tunnel := <-tunnelChan
	defer close(tunnel.Donech)

	s.Write(([]byte)(fmt.Sprintf("LockBox link: http://localhost:3000?id=%d\n", id)))
	s.Write(([]byte)("Session is in progress... Waiting for user to connect and download\n"))
	err := sendFileToTunnel(s.RawCommand(), s, &tunnel)
	if err != nil {
		log.Fatal(err)
		s.Write(([]byte)("Something went wrong!\n"))
		return
	}
	s.Write(([]byte)("File sent successfully, thanks for using LockBox!\n"))
}

func sendFileToTunnel(command string, session io.Reader, tunnel *tunnel.Tunnel) error {
	tunnel.Filename = command
	_, err := io.Copy(tunnel.W, session)
	if err != nil {
		return err
	}
	return nil
}
