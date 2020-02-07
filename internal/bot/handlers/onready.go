package handlers

import (
    "github.com/bwmarrin/discordgo"
)


func OnReady(s *discordgo.Session, event *discordgo.Ready) {
    //s.UpdateStatus(0, "OJ Security")

    log.Infof("%s#%s", event.User.Username, event.User.Discriminator)
}
