package handlers

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func (h *handler) OnMemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	createdAt, _ := discordgo.SnowflakeTimestamp(m.User.ID)

	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	diff := now.Sub(createdAt)

	days := int(diff.Hours()) / 24

	if days >= 2 {
		s.GuildMemberRoleAdd(m.Member.GuildID, m.User.ID, h.Config.Discord.JoinRoleID)
	} else {
		dmchannel, err := s.UserChannelCreate(m.User.ID)
		if err != nil {
			return
		}

		s.ChannelMessageSend(dmchannel.ID, "Sorry your account must be older than 7 days to join, if you believe this is an error please contact an admin.")
	}
}
