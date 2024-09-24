package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var Cfgs Config

const configurations = "./config.yml"

type Config struct {
	Services []Service `yaml:"services"`
}

type Service struct {
	Domain     string   `yaml:"domain"`
	TargetURLs []string `yaml:"targetURLs"`
	Algorythm  string   `yaml:"algorythm"`
}

func GetConfigs() {
	f, err := os.Open(configurations)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfgs)
	if err != nil {
		panic(err)
	}

}
