package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

var config *Config

func Init() {
	//todo multi env
	data, err := os.ReadFile("config/local.yaml")
	if err != nil {
		panic(fmt.Errorf("read config file error: %v", err))
	}

	config = &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		panic(fmt.Errorf("unmarshal config data error: %v", err))
	}
}

func GetConfig() *Config {
	return config
}
