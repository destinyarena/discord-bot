package handlers

import "github.com/bwmarrin/discordgo"

func OnMessageReactionRemove(s *discordgo.Session, mr *discordgo.MessageReactionRemove) {
    reactionroles(s, mr.MessageReaction, false)
}
