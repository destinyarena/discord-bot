package bot

import (
	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
	"github.com/sirupsen/logrus"
)

type (
	// Bot implements gommand router
	Bot struct {
		*gommand.Router
		Config *Config
		Logger *logrus.Logger
		Client *disgord.Client
	}
)

func New(logger *logrus.Logger, client *disgord.Client, config *Config) (*Bot, error) {
	router := gommand.NewRouter(&gommand.RouterConfig{
		PrefixCheck: gommand.MultiplePrefixCheckers(gommand.StaticPrefix(config.Prefix), gommand.MentionPrefix),
	})

	b := &Bot{
		Router: router,
		Config: config,
		Logger: logger,
		Client: client,
	}

	return b, nil
}
