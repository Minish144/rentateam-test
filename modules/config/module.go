package config

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

func New() (*Config, error) {
	m := &Config{
		log: logrus.New(),
	}

	m.log.Infoln("Loading config..")

	if err := m.initalizeConfig(); err != nil {
		m.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Errorf("Can't load config for first time")
		return nil, err
	}
	m.watchChanges()

	m.log.Infoln("Config loaded successfully")

	return m, nil
}
