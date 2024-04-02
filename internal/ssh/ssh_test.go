package ssh

import (
	"bytes"
	"testing"

	"github.com/R894/lockbox/internal/tunnel"
	"github.com/stretchr/testify/assert"
)

func TestSendFileToTunnel(t *testing.T) {
	sessionData := []byte("Test data")
	sessionReader := bytes.NewReader(sessionData)

	tunnel := &tunnel.Tunnel{
		W: &bytes.Buffer{},
	}

	err := sendFileToTunnel("test.txt", sessionReader, tunnel)
	assert.NoError(t, err)

	assert.Equal(t, sessionData, tunnel.W.(*bytes.Buffer).Bytes())
	assert.Equal(t, "test.txt", tunnel.Filename)
}
