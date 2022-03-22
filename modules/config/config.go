package config

import (
	"github.com/spf13/viper"
)

// initializeConfig initialized viper config using yaml config
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
