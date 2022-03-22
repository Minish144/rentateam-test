package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (c *Config) initalizeConfig() error {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil
		}
		return err
	}

	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}

func (c *Config) watchChanges() {
	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		c.log.WithFields(logrus.Fields{
			"name": in.Name,
		}).Info("Config file changed")
	})
}
