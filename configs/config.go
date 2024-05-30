package configs

import (
	"fmt"
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

func GetBaseURL() string {
	return fmt.Sprintf("http://%s:%d/v1/urlshorter",
		Cfg.Service.Address,
		Cfg.Service.Port)
}
