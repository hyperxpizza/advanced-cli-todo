package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		path string `yaml:"path"`
	} `yaml:"database"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}