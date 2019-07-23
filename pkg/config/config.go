package config

import (
	"github.com/socketfunc/colony/proto/config"
	"gopkg.in/yaml.v2"
)

func Parse(file []byte) (*config.Config, error) {
	cfg := new(config.Config)
	err := yaml.Unmarshal(file, cfg)
	return cfg, err
}

type Functions []*config.Function

func (f Functions) Get(name string) *config.Function {
	for i := range f {
		if f[i].Name == name {
			return f[i]
		}
	}
	return nil
}
