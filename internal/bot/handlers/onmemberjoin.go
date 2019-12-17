package handlers

import (
    "github.com/bwmarrin/discordgo"
    "github.com/arturoguerra/d2arena/internal/checkmark"
    "time"
)

func OnMemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
    createdAt, _ := discordgo.SnowflakeTimestamp(m.User.ID)

    loc, _ := time.LoadLocation("UTC")
    now := time.Now().In(loc)
    diff := now.Sub(createdAt)

    days := int(diff.Hours()) / 24

    if (days >= 7) {
        checkmark.Member(s, m.Member)
    } else {
        dmchannel, err := s.UserChannelCreate(m.User.ID)
        if err != nil {
            return
        }

        //s.ChannelMessageSend(dmchannel.ID, "Sorry your account must be older than 7 days to join, if you believe this is an error please contact an admin.")
        s.ChannelMessageSend(dmchannel.ID, "SORRY YOUR ACCOUNT MUST BE OLDER THAN 7 DAYS TO JOIN, IF YOU BELIEVE THIS IS AN ERROR PLEASE CONTACT AN ADMIN.")
    }
}
