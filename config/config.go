package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// MinioConfig holds the configuration for MinIO client.
type MinioConfig struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey"`
	SSL             bool   `mapstructure:"ssl"`
}

// AppConfig holds the application configuration.
type AppConfig struct {
	App struct {
		Name string `mapstructure:"name"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"app"`
	Namespace string      `mapstructure:"namespace"`
	Owner     string      `mapstructure:"owner"`
	Workers   int         `mapstructure:"workers"`
	Minio     MinioConfig `mapstructure:"minio"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (*AppConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // Look for config in the working directory.
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("unable to load configuration: %w", err)
	}
	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to unmarshal viper config: %w", err)
	}
	return &cfg, nil
}
