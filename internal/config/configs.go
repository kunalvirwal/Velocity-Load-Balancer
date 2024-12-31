package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var Cfgs Config

const configurations = "./config.yml"

type Config struct {
	API_PORT    int       `yaml:"api_port"`
	Listen_PORT int       `yaml:"listen_port"`
	Services    []Service `yaml:"services"`
}

type Service struct {
	Domain     string   `yaml:"domain"`
	TargetURLs []string `yaml:"targetURLs"`
	Algorythm  string   `yaml:"algorythm"`
}

// GetConfigs reads the configurations from the config.yml file
func GetConfigs() { // TODO : Implement global error handeling for invalid yaml
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
