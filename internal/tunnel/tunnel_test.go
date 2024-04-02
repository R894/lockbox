package tunnel_test

import (
	"testing"

	"github.com/R894/lockbox/internal/tunnel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTunnelManager(t *testing.T) {
	tm := tunnel.NewTunnelManager()
	require.NotNil(t, tm)
}

func TestAddTunnel(t *testing.T) {
	tm := tunnel.NewTunnelManager()
	id := 1
	tunnelChan := tm.AddTunnel(id)
	assert.NotNil(t, tunnelChan)
}

func TestGetTunnel(t *testing.T) {
	tm := tunnel.NewTunnelManager()
	id := 1
	tunnelChan := tm.AddTunnel(id)
	tunnel, ok := tm.GetTunnel(id)
	assert.True(t, ok)
	assert.Equal(t, tunnelChan, tunnel)
}

func TestInvalidGetTunnel(t *testing.T) {
	tm := tunnel.NewTunnelManager()
	id := 1
	_, ok := tm.GetTunnel(id)
	assert.False(t, ok)
}

func TestRemoveTunnel(t *testing.T) {
	tm := tunnel.NewTunnelManager()
	id := 1
	tm.AddTunnel(id)
	tm.RemoveTunnel(id)
	_, ok := tm.GetTunnel(id)
	assert.False(t, ok)
}
