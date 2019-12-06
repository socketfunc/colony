package api

import (
	"fmt"
	"testing"
)

func TestServer_generatePrivateKey(t *testing.T) {
	server := &server{}
	pk, err := server.generatePrivateKey(4096)
	fmt.Println(err)
	fmt.Println(pk)
}
