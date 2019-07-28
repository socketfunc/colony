package storage

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStorage_Download(t *testing.T) {
	path := "./testdata/test"
	path, err := filepath.Abs(path)
	require.NoError(t, err)

	ctx := context.Background()
	storage := New()
	start := time.Now()
	file, err := storage.Download(ctx, "file://"+path)
	require.NoError(t, err)
	fmt.Println(file)
	fmt.Println(time.Now().Sub(start))

	start = time.Now()
	file, err = storage.Download(ctx, "file://"+path)
	require.NoError(t, err)
	fmt.Println(file)
	fmt.Println(time.Now().Sub(start))
}
