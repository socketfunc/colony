package config

import "sync"

type Store struct {
	data sync.Map
}

func (s *Store) Get(key string) *Config {
	value, ok := s.data.Load(key)
	if !ok {
		return nil
	}
	return value.(*Config)
}

func (s *Store) Save(key string, config *Config) {
	s.data.Store(key, config)
}
