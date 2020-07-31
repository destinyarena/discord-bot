package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func (h *handler) OnReady(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Arena Security")
	h.Logger.Infof("%s#%s", event.User.Username, event.User.Discriminator)
}
