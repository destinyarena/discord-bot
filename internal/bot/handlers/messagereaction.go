package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func hasrole(roles []string, role string) bool {
	for _, i := range roles {
		if i == role {
			return true
		}
	}

	return false
}

func (h *handler) reactionroles(s *discordgo.Session, mr *discordgo.MessageReaction, addrole bool) {

	guild, err := s.Guild(h.Config.Discord.GuildID)
	if err != nil {
		return
	}

	for _, rr := range h.Config.Discord.ReactionRoles {
		if rr.EmojiID == mr.Emoji.APIName() && rr.MessageID == mr.MessageID {
			if addrole {
				s.GuildMemberRoleAdd(guild.ID, mr.UserID, rr.RoleID)
			} else {
				s.GuildMemberRoleRemove(guild.ID, mr.UserID, rr.RoleID)
			}
		}
	}
}
