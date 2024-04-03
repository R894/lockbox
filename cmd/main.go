package main

import (
	"log"

	config "github.com/R894/lockbox"
	"github.com/R894/lockbox/internal/ssh"
	"github.com/R894/lockbox/internal/tunnel"
	"github.com/R894/lockbox/internal/web"
)

func main() {
	cfg := config.LoadConfig()
	tm := tunnel.NewTunnelManager()

	go func() { web.StartServer(cfg.Web.Addr, tm) }()
	log.Fatal(ssh.NewServer(cfg.SSH.Addr, tm).ListenAndServe())
}
