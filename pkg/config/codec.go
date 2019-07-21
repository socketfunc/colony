package config

import "time"

type Config struct {
	Version  string   `yaml:"version"`
	Metadata Metadata `yaml:"metadata"`
	Spec     *Spec    `yaml:"spec"`
}

type Metadata map[string]string

func (m Metadata) Namespace() string {
	return m["namespace"]
}

type Spec struct {
	Functions []*Function `yaml:"functions"`
}

func (s *Spec) GetFunction(name string) *Function {
	for i := range s.Functions {
		if name == s.Functions[i].Name {
			return s.Functions[i]
		}
	}
	return nil
}

type Function struct {
	Name       string     `yaml:"name"`
	Repository string     `yaml:"repository"`
	Resources  *Resources `yaml:"resources"`
}

type Resources struct {
	Requests *Requests `yaml:"requests"`
	Limits   *Limits   `yaml:"limits"`
}

type Requests struct {
	TimeoutStr string `yaml:"timeout"`
}

func (r *Requests) Timeout() (time.Duration, error) {
	return time.ParseDuration(r.TimeoutStr)
}

type Limits struct {
	Requests int `yaml:"requests"`
}
