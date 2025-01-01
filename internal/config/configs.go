package config

import (
	"os"

	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/utils"
	"gopkg.in/yaml.v3"
)

var Cfgs Config

const configurations = "./config.yml"

type Service struct {
	Domain     string   `yaml:"domain"`
	TargetURLs []string `yaml:"targetURLs"`
	Algorythm  string   `yaml:"algorythm"`
}

type Config struct {
	API_PORT    int       `yaml:"api_port"`
	Listen_PORT int       `yaml:"listen_port"`
	Services    []Service `yaml:"services"`
}

// GetConfigs reads the configurations from the config.yml file
func GetConfigs() { // TODO : Implement global error handeling for invalid yaml
	f, err := os.Open(configurations)
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfgs)
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}

}
