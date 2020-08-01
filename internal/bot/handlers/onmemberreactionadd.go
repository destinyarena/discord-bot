package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func (h *handler) OnMessageReactionAdd(s *discordgo.Session, mr *discordgo.MessageReactionAdd) {
	h.reactionroles(s, mr.MessageReaction, true)
}
