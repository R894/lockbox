package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/R894/lockbox/internal/tunnel"
)

func handleRequest(w http.ResponseWriter, r *http.Request, tm *tunnel.TunnelManager) {
	w.Header().Set("Content-Disposition", "attachment; filename=lockbox.bin")
	w.Header().Set("Content-Type", "application/octet-stream")

	idstr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idstr)

	currentTunnel, ok := tm.GetTunnel(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	donech := make(chan struct{})
	currentTunnel <- tunnel.Tunnel{
		W:      w,
		Donech: donech,
	}
	<-donech
}

func StartServer(tm *tunnel.TunnelManager) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, tm)
	})
	fmt.Println("http server is ready")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
