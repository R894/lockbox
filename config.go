package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type webserver struct {
	Host string
	Port int
	Addr string
}
type ssh struct {
	Host    string
	Port    int
	Addr    string
	KeyPath string
}

type Config struct {
	Web webserver
	SSH ssh
}

func LoadConfig() *Config {
	viperInit()
	return &Config{
		Web: webserver{
			Host: viper.GetString("web.host"),
			Port: viper.GetInt("web.port"),
			Addr: fmt.Sprintf("%s:%d", viper.GetString("web.host"), viper.GetInt("web.port")),
		},
		SSH: ssh{
			Host:    viper.GetString("ssh.host"),
			Port:    viper.GetInt("ssh.port"),
			Addr:    fmt.Sprintf("%s:%d", viper.GetString("ssh.host"), viper.GetInt("ssh.port")),
			KeyPath: viper.GetString("ssh.keypath"),
		},
	}
}

func viperInit() {
	setConfigProperties()
	setViperDefaultValues()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file", err)
	}
}

func setViperDefaultValues() {
	viper.SetDefault("web.host", "localhost")
	viper.SetDefault("web.port", 3000)
	viper.SetDefault("ssh.host", "localhost")
	viper.SetDefault("ssh.port", 2222)
}

func setConfigProperties() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}
