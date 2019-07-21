package config

import "gopkg.in/yaml.v2"

func Parse(file []byte) (*Config, error) {
	config := new(Config)
	err := yaml.Unmarshal(file, config)
	return config, err
}
