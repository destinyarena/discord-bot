package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config contains parent config data structure
type Config struct {
	Discord *Discord `yaml:"discord"`
	Faceit  *Faceit  `yaml:"faceit"`
	GRPC    *GRPC    `yaml:"grpc"`
	API     *API     `yaml:"api"`
	NATS    *NATS    `yaml:"nats"`
}

// LoadConfig loads the config for the discord bot
func LoadConfig() (*Config, error) {
	cfgLocation := os.Getenv("DISCORD_CONFIG_LOCATION")

	cfg := Config{}

	if len(cfgLocation) == 0 {
		cfgLocation = "/config/config.yaml"
	}

	yamlFile, err := ioutil.ReadFile(cfgLocation)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(yamlFile, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
