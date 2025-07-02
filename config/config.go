package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Name string `mapstructure:"name"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"app"`
	namespace string `mapstructure:"namespace"`
	owner     string `mapstructure:"owner"`
}

func LoadConfig() (*AppConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("unable to load configuration %v", err))
	}
	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to unmarshal viper config %v", err)
	}
	return &cfg, nil
}
