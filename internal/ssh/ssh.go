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

func StartServer(tm *tunnel.TunnelManager) {
	ssh.Handle(func(s ssh.Session) {
		handleRequest(s, tm)
	})
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}

func handleRequest(s ssh.Session, tm *tunnel.TunnelManager) {
	id := rand.Intn(math.MaxInt)
	tunnelChan := tm.AddTunnel(id)

	s.Write(([]byte)(fmt.Sprintf("LockBox link: http://localhost:3000?id=%d\n", id)))
	fmt.Println("tunnel is ready")
	tunnel := <-tunnelChan
	tunnel.Filename = s.RawCommand()
	_, err := io.Copy(tunnel.W, s)
	if err != nil {
		log.Fatal(err)
	}
	close(tunnel.Donech)
	s.Write(([]byte)("File sent successfully, thanks for using LockBox!\n"))
}
