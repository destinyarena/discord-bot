package handlers

import (
    "github.com/bwmarrin/discordgo"
    "fmt"
)


func OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.ChannelID != "650159281575821312" {
        return
    }

    if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
        fmt.Println("Error deleting message")
    }
}
