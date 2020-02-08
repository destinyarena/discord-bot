package natsevents

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/arturoguerra/d2arena/internal/logging"
    "github.com/bwmarrin/discordgo"
    "fmt"
)

var log = logging.New()
var dcfg = config.LoadDiscord()

func register(s *discordgo.Session, id string) {
    if _, err := s.Guild(dcfg.GuildID); err != nil {
        log.Error(err)
        return
    }

    u, err := s.User(id)
    if err != nil {
        log.Error(err)
        return
    }

    channel, err := s.UserChannelCreate(id)
    if err != nil {
        log.Error(err)
        return
    }

    embed := &discordgo.MessageEmbed{
        Description: "Please click this [link](https://discordapp.com/channels/650109209610027034/657733307353792524/665277347909599232) to get your hub invites!",
    }

    lembed := &discordgo.MessageEmbed{
        Title: "Hub invite link",
        Description: fmt.Sprintf("Sent hubs channel link to <@%s>(`%s#%s`)", id, u.Username, u.Discriminator),
    }

    //All checks are done stuff starts here

    if _, err := s.ChannelMessageSendEmbed(channel.ID, embed); err == nil {
        s.GuildMemberRoleAdd(dcfg.GuildID, id, dcfg.RegistrationRoleID)
    } else {
        lembed = &discordgo.MessageEmbed{
            Title: "403: Forbidden",
            Description: fmt.Sprintf("Error sending hub channel to <@%s>(`%s#%s`) please contact them", id, u.Username, u.Discriminator),
        }
    }

    s.ChannelMessageSendEmbed(dcfg.LogsID, lembed)
}

func registration(dg *discordgo.Session, nchan *structs.NATS) {
    for i := range nchan.RecvRegistration {
        if i.Id != "" {
            log.Infof("Registering user: %s", i.Id)
            register(dg, i.Id)
        }
    }
}


func New(dg *discordgo.Session, nchan *structs.NATS) {
    log.Infoln("Registering NATS Events")
    go registration(dg, nchan)
}

