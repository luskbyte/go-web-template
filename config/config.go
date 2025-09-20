package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	ServerConfig `mapstructure:"server_config"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

func LoadConfig() (*config, error) {
	cfg := &config{}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("Fatal error unmarshal config file: %w \n", err)
	}

	return cfg, nil
}
