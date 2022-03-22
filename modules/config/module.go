package config

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

func New() (*Config, error) {
	m := &Config{}

	if err := m.initalizeConfig(); err != nil {
		return nil, err
	}

	return m, nil
}
