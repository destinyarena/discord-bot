package handlers

import (
    "context"
    pb "github.com/arturoguerra/d2arena/pkg/faceit/proto"
    //"github.com/arturoguerra/d2arena/internal/structs"
    "google.golang.org/grpc"

//    "github.com/arturoguerra/d2arena/pkg/faceit"
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

func testrpc(id string) {
    grpcfg := config.LoadgRPC()
    address := fmt.Sprintf("%s:%s", grpcfg.FaceitHost, grpcfg.FaceitPort)
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()
    fmt.Println("Connected again")

    c := pb.NewFaceitClient(conn)

    r, err := c.GetProfile(context.Background(), &pb.ProfileRequest{
        Guid: id,
    })

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("%s %s %s", r.GetGuid(), r.GetSkilllvl(), r.GetUsername())
}


func getInvite(hubid string) (string, error) {
    grpcfg := config.LoadgRPC()
    address := fmt.Sprintf("%s:%s", grpcfg.FaceitHost, grpcfg.FaceitPort)
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        fmt.Println(err)
        return "", err
    }
    defer conn.Close()
    fmt.Println("Connected")

    c := pb.NewFaceitClient(conn)

    r, err := c.GetInvite(context.Background(), &pb.InviteRequest{
        Hubid: hubid,
    })
    if err != nil {
        fmt.Println(err)
        return "", err
    }

    link := fmt.Sprintf("%s/%s", r.GetBase(), r.GetCode())
    return link, nil
}

func fetchProfile(id string) (*Profile, error) {
    go testrpc(id)
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

    if resp.StatusCode != 200 && resp.StatusCode != 201 {
        err = fmt.Errorf("Server returned code: %d", resp.StatusCode)
        return nil, err
    }

    rawbody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var body Profile
    json.Unmarshal([]byte(rawbody), &body)
    v := validator.New()
    if err = v.Struct(body); err != nil {
        return nil, err
    }

    return &body, nil
}


func invites(s *discordgo.Session, mr *discordgo.MessageReactionAdd) {
    title := "Destiny Arena Faceit Invitation"

    guild, err := s.Guild(cfg.GuildID)
    if err != nil {
        return
    }

    u, err := s.User(mr.UserID)
    if err != nil {
        return
    }

    channel, err := s.UserChannelCreate(mr.UserID)
    if err != nil {
        return
    }

    member, err := getMember(guild, mr.UserID)
    if err != nil {
        return
    }

    var mainhubs string
    var addithubs string
    send := false
    var roles []string
    var profile *Profile


    for _, hub := range cfg.Hubs {
        if hub.EmojiID == mr.Emoji.APIName() && hub.MessageID == mr.MessageID {
            if profile == nil {
                profile, err = fetchProfile(mr.UserID)
                if err != nil {
                    fmt.Println(err)
                    embed := &discordgo.MessageEmbed{
                        Title: title,
                        Description: "Looks like you haven't Registered yet, please do that before requesting invites.",
                    }

                    s.ChannelMessageSendEmbed(channel.ID, embed)
                    return
                }
            }

            if checkRole(member.Roles, hub.RoleID) {
                fmt.Println("-----------")
                fmt.Println(profile.Faceitlvl)
                fmt.Println(hub.SkillLvl)
                if profile.Faceitlvl >= hub.SkillLvl {
                    fmt.Println("Getting invite..")
                    link, err := getInvite(hub.HubID)
                    if err != nil || link == "" {
                        fmt.Println(err)
                    } else {
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

func OnMessageReactionAdd(s *discordgo.Session, mr *discordgo.MessageReactionAdd) {
    invites(s, mr)
    reactionroles(s, mr.MessageReaction, true)
}
