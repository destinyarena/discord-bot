package handlers

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
)


func OnReady(s *discordgo.Session, event *discordgo.Ready) {
    s.UpdateStatus(0, "OJ Security")

    fmt.Printf("%s#%s", event.User.Username, event.User.Discriminator)
}
