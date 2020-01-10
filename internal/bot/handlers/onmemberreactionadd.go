package handlers

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "github.com/arturoguerra/d2arena/internal/config"
    "gopkg.in/go-playground/validator.v9"
    "github.com/bwmarrin/discordgo"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "errors"
    "fmt"
)

var cfg *structs.Discord

type Profile struct {
    DiscordID string `json:"discordid" validate:"required"`
    SteamID string `json:"steamid" validate:"required"`
    FaceitGuid string `json:"faceitguid" validate:"required"`
    FaceitName string `json:"faceitname" valitate:"required"`
}

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


func fetchDB(id string) (*Profile, error) {
    base := fmt.Sprintf("https://destinyarena.fireteamsupport.net/infoexchange.php?key=2YHSbPt5GJ9Uupgk&d=true&discordid=%s", id)

    client := new(http.Client)
    req, _ := http.NewRequest("GET", base, nil)
    resp, err := client.Do(req)

    if err != nil {
        return nil, err
    }

    rawbody, _ := ioutil.ReadAll(resp.Body)
    var body Profile
    json.Unmarshal([]byte(rawbody), &body)
    v := validator.New()
    if err = v.Struct(body); err != nil {
        return nil, err
    }

    return &body, nil
}

type GameBody struct {
    SkillLevel int `json:"skill_level" validate:"required"`
}

type ReqBody struct {
    Games map[string]GameBody `json:"games" validate:"required"`
}

func getFaceitLevel(userid string) int {
    profile, err := fetchDB(userid)
    if err != nil {
        fmt.Println("Error fetching profile")
        return 0
    }

    // FaceitGuid
    config := config.LoadFaceit()
    client := new(http.Client)
    url := "https://open.faceit.com/data/v4/players/" + profile.FaceitGuid
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Add("Authorization", "Bearer " + config.ApiToken)
    req.Header.Add("Content-Type", "application/json")
    resp, err := client.Do(req)
    defer resp.Body.Close()

    if err != nil || resp.StatusCode != 200 {
        fmt.Println("Error fetching faceit profile")
        return 0
    }

    rawbody, _ := ioutil.ReadAll(resp.Body)

    var body ReqBody
    json.Unmarshal([]byte(rawbody), &body)
    v := validator.New()
    if err = v.Struct(body); err != nil {
        fmt.Println(err)
        return 0
    }

    if val, ok := body.Games["destiny2"]; ok {
        fmt.Println(val)
        return val.SkillLevel
    }

    return 0
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

    faceitlevel := getFaceitLevel(mr.UserID)

    for _, hub := range cfg.Hubs {
        if hub.EmojiID == mr.Emoji.APIName() || mr.Emoji.APIName() == cfg.InvitesAutoEmojiID {
            if checkRole(member.Roles, hub.RoleID) {
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
    if cfg.InvitesMessageID == mr.MessageID {
        invites(s, mr)
    } else if cfg.RulesMessageID == mr.MessageID {
        rules(s, mr)
    }
}
