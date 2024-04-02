package tunnel

import (
	"io"
	"sync"
)

type Tunnel struct {
	Filename string
	W        io.Writer
	Donech   chan struct{}
}

type TunnelManager struct {
	tunnels map[int]chan Tunnel
	mu      sync.Mutex
}

func NewTunnelManager() *TunnelManager {
	return &TunnelManager{
		tunnels: make(map[int]chan Tunnel),
	}
}

func (tm *TunnelManager) AddTunnel(id int) chan Tunnel {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.tunnels[id] = make(chan Tunnel)
	return tm.tunnels[id]
}

func (tm *TunnelManager) GetTunnel(id int) (chan Tunnel, bool) {
	tunnel, ok := tm.tunnels[id]
	return tunnel, ok
}

func (tm *TunnelManager) RemoveTunnel(id int) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	close(tm.tunnels[id])
	delete(tm.tunnels, id)
}
