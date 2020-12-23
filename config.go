package main

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// Config documented in Readme https://yaml.to-go.online/
type Config struct {
	Domain        string `yaml:"domain"`
	Port          string `yaml:"port"`
	StorageEngine string `yaml:"storageEngine"`
	Localfile     struct {
		BasePath string `yaml:"basePath"`
	} `yaml:"localfile"`
}

var config Config

func parseConfig() {
	dataFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(dataFile, &config)
}
