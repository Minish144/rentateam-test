package config

import (
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
