package storage

//go:generate mkdir -p mock
//go:generate mockgen -source=storage.go -package=mock -destination=mock/storage.go
import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/ReneKroon/ttlcache"
	errors "golang.org/x/xerrors"
)

const (
	fileProtocol  = "file://"
	httpProtocol  = "http://"
	httpsProtocol = "https://"
)

var (
	expiration = time.Duration(5 * time.Minute)
)

type Storage interface {
	Download(ctx context.Context, path string) ([]byte, error)
}

type storage struct {
	cache *ttlcache.Cache
}

func New() Storage {
	return &storage{
		cache: ttlcache.NewCache(),
	}
}

func (s *storage) Download(ctx context.Context, path string) ([]byte, error) {
	if file, ok := s.cache.Get(path); ok {
		return file.([]byte), nil
	}
	var (
		file []byte
		err  error
	)
	switch {
	case strings.HasPrefix(path, fileProtocol):
		file, err = s.getLocalFile(path)
	case strings.HasPrefix(path, httpProtocol), strings.HasPrefix(path, httpsProtocol):
		file, err = s.getHttpFile(ctx, path)
	default:
		err = errors.New("storage: unsupported storage protocol")
	}
	if err != nil {
		return nil, err
	}
	s.cache.SetWithTTL(path, file, expiration)
	return file, err
}

func (s *storage) getLocalFile(path string) ([]byte, error) {
	path = strings.TrimPrefix(path, fileProtocol)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *storage) getHttpFile(ctx context.Context, path string) ([]byte, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
