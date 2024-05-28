package config

import (
	"go-short-url/utils"
	"os"
)

const (
	DefaultConfigFileName string = "default.yaml"
	DefaultConfigFolder   string = "./defaults"
)

var Cfg *Config

type Config struct {
	Service struct {
		Port    int    `yaml:"port"`
		Address string `yaml:"address"`
	} `yaml:"service"`
}

func ReadConfig() {
	Cfg = &Config{}
	defaultConfigFile := DefaultConfigFolder + "/" + DefaultConfigFileName
	if err := utils.ParseYamlFile(defaultConfigFile, Cfg); err != nil {
		os.Exit(2)
	}
}
