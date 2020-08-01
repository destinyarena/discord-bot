package handlers

import (
	"github.com/arturoguerra/d2arena/internal/config"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type (
	handler struct {
		Config *config.Config
		Logger *logrus.Logger
	}

	// Handler exports discordgo event handlers
	Handler interface {
		OnReady(*discordgo.Session, *discordgo.Ready)
		OnMessageReactionAdd(*discordgo.Session, *discordgo.MessageReactionAdd)
		OnMessageReactionRemove(*discordgo.Session, *discordgo.MessageReactionRemove)
	}
)

// New returns a handler interface
func New(cfg *config.Config, log *logrus.Logger) Handler {
	return &handler{
		Config: cfg,
		Logger: log,
	}
}
