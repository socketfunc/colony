package storage

import (
	"io/ioutil"
	"strings"

	errors "golang.org/x/xerrors"
)

const (
	fileProtocol = "file://"
)

type Storage interface {
	Download(path string) ([]byte, error)
}

type storage struct {
}

func New() Storage {
	return &storage{}
}

func (s *storage) Download(path string) ([]byte, error) {
	switch {
	case strings.HasPrefix(path, fileProtocol):
		return s.local(path)
	}
	return nil, errors.New("storage:")
}

func (s *storage) local(path string) ([]byte, error) {
	path = strings.TrimPrefix(path, fileProtocol)
	return ioutil.ReadFile(path)
}
