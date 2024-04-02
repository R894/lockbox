package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port int
}

func LoadConfig() *Config {
	viperInit()
	return &Config{
		Host: viper.GetString("host"),
		Port: viper.GetInt("port"),
	}
}

func viperInit() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("host", "0.0.0.0")
	viper.SetDefault("port", 2222)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file", err)
	}
}
