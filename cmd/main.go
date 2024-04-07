package main

import (
	"log"

	config "github.com/R894/lockbox"
	"github.com/R894/lockbox/internal/ssh"
	"github.com/R894/lockbox/internal/tunnel"
	"github.com/R894/lockbox/internal/web"
)

func main() {
	var (
		cfg           *config.Config        = config.LoadConfig()
		tunnelManager *tunnel.TunnelManager = tunnel.NewTunnelManager()
		webAddr       string                = cfg.Web.Addr
		sshAddr       string                = cfg.SSH.Addr
		keyPath       string                = cfg.SSH.KeyPath
	)

	go func() { web.StartServer(webAddr, tunnelManager) }()
	log.Fatal(ssh.NewServer(sshAddr, tunnelManager, keyPath).ListenAndServe())
}
