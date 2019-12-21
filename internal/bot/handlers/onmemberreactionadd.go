package handlers

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/bwmarrin/discordgo"
    "strings"
    "errors"
    "fmt"
)

var cfg *structs.Discord

func init() {
    cfg = config.LoadDiscord()
}

func checkRole(roles []string, role string) bool {
    for _, r := range roles {
        if r == role {
            return false
        }
    }

    return true
}

func getMember(g *discordgo.Guild, uid string) (*discordgo.Member, error) {
    for _, member := range g.Members {
        if member.User.ID == uid {
            return member, nil
        }
    }

    return nil, errors.New("Not found")
}

func invites(s *discordgo.Session, mr *discordgo.MessageReactionAdd) {
    guild, err := s.Guild(cfg.GuildID)
    if err != nil {
        return
    }

    member, err := getMember(guild, mr.UserID)
    if err != nil {
        return
    }

    message := "Invite links to join FACEIT Hubs\n\n"
    send := false
    var roles []string

    for _, reaction := range cfg.Reactions {
        if reaction.EmojiID == mr.Emoji.APIName() || mr.Emoji.APIName() == cfg.InvitesAutoEmojiID {
            if checkRole(member.Roles, reaction.RoleID) {
                if link, _ := getInvite(reaction.HubID); link != "" {
                    roles = append(roles, reaction.RoleID)
                    message += strings.Replace(reaction.Format, "{invite}", link, 1) + "\n"
                    send = true
                }
            }
        }
    }

    if send {
        channel, _ := s.UserChannelCreate(mr.UserID)
        if _, err := s.ChannelMessageSend(channel.ID, message); err == nil {
            for _, role := range roles {
                s.GuildMemberRoleAdd(cfg.GuildID, mr.UserID, role)
            }
        }
    }
}

func rules(s *discordgo.Session, mr *discordgo.MessageReactionAdd) {
    if mr.Emoji.APIName() != cfg.RulesEmojiID {
        return
    }

    s.GuildMemberRoleAdd(cfg.GuildID, mr.UserID, cfg.RulesRoleID)



}


func OnMessageReactionAdd(s *discordgo.Session, mr *discordgo.MessageReactionAdd) {
    fmt.Println(mr.Emoji.APIName())
    fmt.Println(cfg.InvitesAutoEmojiID)
    fmt.Println(mr.Emoji.APIName() == cfg.InvitesAutoEmojiID)
    if cfg.InvitesMessageID == mr.MessageID {
        invites(s, mr)
    } else if cfg.RulesMessageID == mr.MessageID {
        rules(s, mr)
    }
}
