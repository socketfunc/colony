package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGCM(t *testing.T) {
	key := []byte("mnSijgCt4ZnxRYB8")
	gcm := NewGCM()
	assert.Equal(t, GCM, gcm.Type())
	data := []byte("text...")
	enc, err := gcm.Encrypt(key, data)
	require.NoError(t, err)
	src, err := gcm.Decrypt(key, enc)
	require.NoError(t, err)
	assert.Equal(t, data, src)
}
