package checkmark

import (
    "github.com/bwmarrin/discordgo"
    "github.com/arturoguerra/d2arena/internal/config"
)
func Member(s *discordgo.Session, m *discordgo.Member) {
    discord := config.LoadDiscord()

    s.GuildMemberRoleAdd(m.GuildID, m.User.ID, discord.JoinRoleID)
    s.GuildMemberRoleAdd(m.GuildID, m.User.ID, discord.DividerRoleID)
}
