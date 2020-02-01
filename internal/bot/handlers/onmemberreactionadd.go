package handlers

import (
    //"github.com/arturoguerra/d2arena/internal/structs"
    "github.com/arturoguerra/d2arena/internal/config"
    "gopkg.in/go-playground/validator.v9"
    "github.com/bwmarrin/discordgo"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "errors"
    "fmt"
)

var cfg = config.LoadDiscord()
var apicfg = config.LoadAPIConfig()

type Profile struct {
    Discord   string `json:"discord"   validate:"required"`
    Bungie    string `json:"bungie"    validate:"required"`
    Faceit    string `json:"faceit"    validate:"required"`
    Faceitlvl int    `json:"faceitlvl" validate:"required"`
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


func fetchProfile(id string) (*Profile, error) {
    token := config.LoadAuth()
    base := fmt.Sprintf("%s/api/users/get/%s", apicfg.BaseURL, id)

    client := new(http.Client)
    req, err := http.NewRequest("GET", base, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer " + token)
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }

    rawbody, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(rawbody))
    var body Profile
    json.Unmarshal([]byte(rawbody), &body)
    v := validator.New()
    if err = v.Struct(body); err != nil {
        return nil, err
    }

    return &body, nil
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

    title := "Destiny Arena Faceit Invitation"
    var mainhubs string
    var addithubs string
    send := false
    var roles []string

    profile, err := fetchProfile(mr.UserID)
    if err != nil {
        fmt.Println(err)
        profile = &Profile{
            Faceitlvl: 0,
        }
    }

    faceitlevel := profile.Faceitlvl

    for _, hub := range cfg.Hubs {
        if hub.EmojiID == mr.Emoji.APIName() && hub.MessageID == mr.MessageID {
            if checkRole(member.Roles, hub.RoleID) {
                fmt.Println("-----------")
                fmt.Println(faceitlevel)
                fmt.Println(hub.SkillLvl)
                if faceitlevel >= hub.SkillLvl {
                    if link, _ := getInvite(hub.HubID); link != "" {
                        roles = append(roles, hub.RoleID)
                        if hub.Main {
                            if mainhubs == "" {
                                mainhubs += "Main Hubs:\n"
                            }

                            mainhubs += fmt.Sprintf("[%s](%s)\n", hub.Format, link)
                        } else {
                            if addithubs == "" {
                                addithubs += "Additional Hubs:\n"
                            }

                            addithubs += fmt.Sprintf("[%s](%s)\n", hub.Format, link)
                        }

                        send = true
                    }
                }
            }
        }
    }

    if send {
        u, err := s.User(mr.UserID)
        if err != nil {
            return
        }
        channel, _ := s.UserChannelCreate(mr.UserID)
        mainhubs += "\n"
        message := mainhubs + addithubs
        embed := &discordgo.MessageEmbed{
            Title: title,
            Description: message,
        }

        if _, err := s.ChannelMessageSendEmbed(channel.ID, embed); err == nil {
            for _, role := range roles {
                s.GuildMemberRoleAdd(cfg.GuildID, mr.UserID, role)
            }
            s.ChannelMessageSendEmbed(cfg.LogsID, &discordgo.MessageEmbed{
                Title: "Notification",
                Description: fmt.Sprintf("User <@%s>(`%s#%s`) requested invites", mr.UserID, u.Username, u.Discriminator),
            })
        } else {
            embed := &discordgo.MessageEmbed{
                Title: "403: Forbidden",
                Description: fmt.Sprintf("Error sending invites to <@%s>(`%s#%s`)", mr.UserID, u.Username, u.Discriminator),
            }

            s.ChannelMessageSendEmbed(cfg.LogsID, embed)
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
    if cfg.RulesMessageID == mr.MessageID {
        rules(s, mr)
    } else {
        invites(s, mr)
    }
}


func phpFix(id string, skilllvl int) {
    base := fmt.Sprintf("http://destinyarena.fireteamsupport.net/updateskill.php?key=2YHSbPt5GJ9Uupgk&f=%s&s=%d", id, skilllvl)

    if _, err := http.Get(base); err != nil {
        fmt.Println(err)
    }
}
