package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/bot/internal/config"
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
		OnMemberJoin(*discordgo.Session, *discordgo.GuildMemberAdd)
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
