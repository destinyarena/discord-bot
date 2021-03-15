package handlers

import (
	"github.com/destinyarena/bot/internal/bot"
	"github.com/destinyarena/bot/internal/faceit"
	"github.com/destinyarena/bot/internal/profiles"
	"github.com/sirupsen/logrus"
)

type (
	handler struct {
		Bot      *bot.Bot
		Faceit   faceit.Faceit
		Profiles profiles.Profiles
		Logger   *logrus.Logger
	}

	// Handler exports discordgo event handlers
//	Handler interface {
//		OnReady(*discordgo.Session, *discordgo.Ready)
//		OnMemberJoin(*discordgo.Session, *discordgo.GuildMemberAdd)
//		OnMessageReactionAdd(*discordgo.Session, *discordgo.MessageReactionAdd)
//		OnMessageReactionRemove(*discordgo.Session, *discordgo.MessageReactionRemove)
//	}
)

// New returns a handler interface
func New(bot *bot.Bot, faceit faceit.Faceit, profiles profiles.Profiles, log *logrus.Logger) {
	h := &handler{
		Bot:      bot,
		Faceit:   faceit,
		Profiles: profiles,
		Logger:   log,
	}

	bot.Client.Gateway().BotReady(h.OnReady())
	bot.Client.Gateway().GuildMemberAdd(h.OnMemberJoin())
	//bot.Client.Gateway().MessageReactionAdd(h.OnMessageReactionAdd())
	//bot.Client.Gateway().MessageReactionRemove(h.OnMessageReactionRemove())

}
