package testhelper

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		ConString string `yaml:"con_string"`
		Dialect   string `yaml:"dialect"`
	} `yaml:"database"`
}

func NewFromFile(filename string) (Config, error) {
	var config Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return config, nil
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, nil
	}

	return config, nil
}
