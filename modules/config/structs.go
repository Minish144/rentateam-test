package config

// AppConfig is an app settings structure
type AppConfig struct {
	LogLevel string `mapstructure:"log_level"`
}

// ServerConfig is a server settings structure
type ServerConfig struct {
	GRPC GRPCConfig `mapstructure:"grpc"`
	HTTP HTTPConfig `mapstructure:"http"`
}

// GRPCConfig is a grpc settings structure
type GRPCConfig struct {
	Interface string `mapstructure:"interface"`
	Port      uint32 `mapstructure:"port"`
}

// HTTPConfig is an http settings structure
type HTTPConfig struct {
	Interface string `mapstructure:"interface"`
	Port      uint32 `mapstructure:"port"`
}

// DatabaseConfig is a postgresql settings structure
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     uint32 `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}
