package main

import (
	"github.com/R894/lockbox/internal/ssh"
	"github.com/R894/lockbox/internal/tunnel"
	"github.com/R894/lockbox/internal/web"
)

func main() {
	tm := tunnel.NewTunnelManager()
	go func() { web.StartServer(tm) }()
	ssh.StartServer(tm)
}
