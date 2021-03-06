package main

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// Config documented in Readme https://yaml.to-go.online/
type Config struct {
	Domain         string `yaml:"domain"`
	Port           string `yaml:"port"`
	LogLevel       string `yaml:"logLevel"`
	StorageEngine  string `yaml:"storageEngine"`
	DropExpiration int    `yaml:"dropExpiration"`
	Localfile      struct {
		BasePath string `yaml:"basePath"`
	} `yaml:"localfile"`
	Redis struct {
		Addr     string      `yaml:"addr"`
		Port     string      `yaml:"port"`
		Password interface{} `yaml:"password"`
	} `yaml:"redis"`
}

var config Config

func parseConfig() {
	dataFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(dataFile, &config)
}
