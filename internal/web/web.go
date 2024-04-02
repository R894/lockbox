package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/R894/lockbox/internal/tunnel"
)

func handleRequest(w http.ResponseWriter, r *http.Request, tm *tunnel.TunnelManager) {
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	currentTunnel, ok := tm.GetTunnel(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	sendResponse(w, currentTunnel)
}

func sendResponse(w http.ResponseWriter, currentTunnel chan tunnel.Tunnel) {
	w.Header().Set("Content-Disposition", "attachment; filename=lockbox.bin")
	w.Header().Set("Content-Type", "application/octet-stream")

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
