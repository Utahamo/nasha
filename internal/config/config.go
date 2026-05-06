package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Auth     AuthConfig     `yaml:"auth"`
	Mounts   []MountConfig  `yaml:"mounts"`
}

type ServerConfig struct {
	Addr       string `yaml:"addr"`
	StaticDir  string `yaml:"static_dir"`
}

type DatabaseConfig struct {
	DSN string `yaml:"dsn"`
}

type AuthConfig struct {
	JWTSecret string `yaml:"jwt_secret"`
	TokenTTL  string `yaml:"token_ttl"`
}

type MountConfig struct {
	Name   string            `yaml:"name"`
	Type   string            `yaml:"type"`
	Path   string            `yaml:"path"`
	Config map[string]string `yaml:"config"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
