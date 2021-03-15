package handlers

//func (h *handler) reactionroles(s *discordgo.Session, mr *discordgo.MessageReaction, addrole bool) {
//	guild, err := s.Guild(h.Config.Discord.GuildID)
//	if err != nil {
//		return
//	}
//
//	for _, rr := range h.Bot.Config.ReactionRoles {
//		emoji := mr.Emoji.ID
//		if len(emoji) == 0 {
//			emoji = mr.Emoji.APIName()
//		}
//
//		if rr.EmojiID == emoji && rr.MessageID == mr.MessageID {
//			if addrole {
//				s.GuildMemberRoleAdd(guild.ID, mr.UserID, rr.RoleID)
//			} else {
//				s.GuildMemberRoleRemove(guild.ID, mr.UserID, rr.RoleID)
//			}
//		}
//	}
//}
