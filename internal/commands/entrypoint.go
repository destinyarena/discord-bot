package commands

import (
	"github.com/destinyarena/bot/internal/bot"
	"github.com/destinyarena/bot/internal/faceit"
	"github.com/destinyarena/bot/internal/profiles"
	"github.com/sirupsen/logrus"
)

type BaseCommand struct {
	Faceit   faceit.Faceit
	Profiles profiles.Profiles
	Config   *bot.Config
	Logger   *logrus.Logger
}

// New returns a new command handler
func New(bot *bot.Bot, faceit faceit.Faceit, profiles profiles.Profiles, logger *logrus.Logger) {
	base := BaseCommand{
		Faceit:   faceit,
		Profiles: profiles,
		Logger:   logger,
	}

	bot.SetCommand(&ban{BaseCommand: base})
	bot.SetCommand(&unban{BaseCommand: base})
	bot.SetCommand(&clear{BaseCommand: base})
	bot.SetCommand(&profile{BaseCommand: base})

}
