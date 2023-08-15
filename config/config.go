package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Postgres `yaml:"postgres"`
	Telebot  `yaml:"telegram"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
	SSLmode  string `yaml:"sslmode"`
}

type Telebot struct {
	Token string `yaml:"token"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	content, err := os.ReadFile(filepath.Join("config", "config.yml"))
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = yaml.Unmarshal(content, cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal data: %w", err)
	}

	return cfg, nil
}
