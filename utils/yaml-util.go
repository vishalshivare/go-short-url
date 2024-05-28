package utils

import (
	"log"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
)

func ParseYamlFile(yamlFileName string, out interface{}) error {
	fileName, err := filepath.Abs(yamlFileName)
	if err != nil {
		log.Printf("error getting absolute path of %s Err: %v", yamlFileName, err)
		return err
	}

	yamlBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("error reading file %s Err: %v", fileName, err)
		return err
	}
	if err := yaml.Unmarshal(yamlBytes, out); err != nil {
		log.Printf("error parsing the yaml file %s Err: %v", fileName, err)
		return err
	}
	return nil
}
