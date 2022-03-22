package config

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

// Config is a config structure
type Config struct {
	log      *logrus.Logger
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

// New creates new object of Config module
func New() (*Config, error) {
	m := &Config{}

	if err := m.initalizeConfig(); err != nil {
		return nil, err
	}

	return m, nil
}
