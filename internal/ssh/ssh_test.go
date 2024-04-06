package ssh

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/R894/lockbox/internal/tunnel"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ssh"
)

func RunCmdOverSSH(addr, username, password, cmd string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cfg := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", addr, cfg)
	if err != nil {
		return nil, err
	}

	session, err := conn.NewSession()
	if err != nil {
		return nil, err
	}
	go func() {
		<-ctx.Done()
		session.Close()
	}()
	return session.CombinedOutput(cmd)
}

func TestFileTransfer(t *testing.T) {
	sessionData := []byte("Test data")
	sessionReader := bytes.NewReader(sessionData)

	tunnel := &tunnel.Tunnel{
		W: &bytes.Buffer{},
	}

	err := startFileTransfer(sessionReader, tunnel)
	assert.NoError(t, err)

	assert.Equal(t, sessionData, tunnel.W.(*bytes.Buffer).Bytes())
}

func TestStartServer(t *testing.T) {
	fmt.Println("IN test")
	tm := tunnel.NewTunnelManager()
	srv := NewServer("127.0.0.1:2222", tm, "")
	go srv.ListenAndServe()

	res, _ := RunCmdOverSSH("127.0.0.1:2222", "root", "root", "test.txt")
	resStr := string(res)
	fmt.Println("got string: ", resStr)
	assert.Contains(t, resStr, "Session is in progress... Waiting for user to connect and download")
}
