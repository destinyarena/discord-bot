package config

import (
	"io/ioutil"
	"os"

	"github.com/destinyarena/bot/internal/bot"
	"github.com/destinyarena/bot/internal/faceit"
	"github.com/destinyarena/bot/internal/natsevents"
	"github.com/destinyarena/bot/internal/profiles"

	"gopkg.in/yaml.v2"
)

// Config contains parent config data structure
type Config struct {
	Discord  *bot.Config        `yaml:"discord"`
	Faceit   *faceit.Config     `yaml:"faceit"`
	Profiles *profiles.Config   `yaml:"profiles"`
	NATS     *natsevents.Config `yaml:"nats"`
}

// LoadConfig loads the config for the discord bot
func LoadConfig() (*Config, error) {
	cfgLocation := os.Getenv("BOT_CONFIG_LOCATION")

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
