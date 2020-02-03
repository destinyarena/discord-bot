package handlers

import (
    "github.com/bwmarrin/discordgo"
    "github.com/arturoguerra/d2arena/internal/config"
)

func hasrole(roles []string, role string) bool {
    for _, i := range roles {
        if i == role {
            return true
        }
    }

    return false
}

func reactionroles(s *discordgo.Session, mr *discordgo.MessageReaction, addrole bool) {
    cfg := config.LoadDiscord()

    guild, err := s.Guild(cfg.GuildID)
    if err != nil {
        return
    }

    for _, rr := range cfg.ReactionRoles {
        if rr.EmojiID == mr.Emoji.APIName() && rr.MessageID == mr.MessageID {
            if addrole {
                s.GuildMemberRoleAdd(guild.ID, mr.UserID, rr.RoleID)
            } else {
                s.GuildMemberRoleRemove(guild.ID, mr.UserID, rr.RoleID)
            }
        }
    }
}
