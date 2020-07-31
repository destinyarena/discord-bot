package handlers

import "github.com/bwmarrin/discordgo"

func (h *handler) OnMessageReactionRemove(s *discordgo.Session, mr *discordgo.MessageReactionRemove) {
	h.reactionroles(s, mr.MessageReaction, false)
}
