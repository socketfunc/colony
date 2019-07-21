package storage

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStorage_Download(t *testing.T) {
	path := "./testdata/test"
	path, err := filepath.Abs(path)
	require.NoError(t, err)

	storage := New()
	file, err := storage.Download("file://" + path)
	require.NoError(t, err)
	fmt.Println(len(file))
}
