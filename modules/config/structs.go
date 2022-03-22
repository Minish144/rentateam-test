package config

import "github.com/sirupsen/logrus"

type Config struct {
	log      *logrus.Logger
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type AppConfig struct {
	LogLevel string `mapstructure:"log_level"`
	SaveLogs bool   `mapstructure:"save_logs"`
}

type ServerConfig struct {
	GRPC GRPCConfig `mapstructure:"grpc"`
	HTTP HTTPConfig `mapstructure:"http"`
}

type GRPCConfig struct {
	Interface string `mapstructure:"interface"`
	Port      uint32 `mapstructure:"port"`
}

type HTTPConfig struct {
	Enabled   bool   `mapstructure:"enabled"`
	Interface string `mapstructure:"interface"`
	Port      uint32 `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     uint32 `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}
