package handlers

import (
    "context"
    profiles "github.com/arturoguerra/d2arena/pkg/profiles"
    faceit "github.com/arturoguerra/d2arena/pkg/faceit"
    "google.golang.org/grpc"

    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/bwmarrin/discordgo"
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


func getInvite(hubid string) (string, error) {
    grpcfg := config.LoadgRPC()
    address := fmt.Sprintf("%s:%s", grpcfg.FaceitHost, grpcfg.FaceitPort)
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Error(err)
        return "", err
    }
    defer conn.Close()

    c := faceit.NewFaceitClient(conn)
    log.Infoln("Fetching Invites")
    r, err := c.GetInvite(context.Background(), &faceit.InviteRequest{
        Hubid: hubid,
    })
    if err != nil {
        log.Error(err)
        return "", err
    }

    link := fmt.Sprintf("%s/%s", r.GetBase(), r.GetCode())
    return link, nil
}

func fetchProfile(id string) (*Profile, error) {
    grpcfg := config.LoadgRPC()
    address := fmt.Sprintf("%s:%s", grpcfg.ProfilesHost, grpcfg.ProfilesPort)
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Error(err)
        return nil, err
    }

    defer conn.Close()

    c := profiles.NewProfilesClient(conn)
    log.Infoln("Fetching database profile")
    r, err := c.GetProfile(context.Background(), &profiles.IdRequest{
        Id: id,
    })

    if err != nil {
        log.Error(err)
        return nil, err
    }

    address = fmt.Sprintf("%s:%s", grpcfg.FaceitHost, grpcfg.FaceitPort)
    conn, err = grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Error(err)
        return nil, err
    }

    defer conn.Close()
    log.Infoln("Fetching faceit skill level")

    f := faceit.NewFaceitClient(conn)
    rf, err := f.GetProfile(context.Background(), &faceit.ProfileRequest{
        Guid: r.GetFaceit(),
    })
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return &Profile{
        Discord: r.GetDiscord(),
        Bungie: r.GetBungie(),
        Faceit: r.GetFaceit(),
        Faceitlvl: int(rf.GetSkilllvl()),
    }, nil
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
                    log.Errorln(err)
                    embed := &discordgo.MessageEmbed{
                        Title: title,
                        Description: "Looks like you haven't Registered yet, please do that before requesting invites.",
                    }

                    s.ChannelMessageSendEmbed(channel.ID, embed)
                    return
                }
            }

            if checkRole(member.Roles, hub.RoleID) {
                log.Infoln("-----------")
                log.Infoln(profile.Faceitlvl)
                log.Infoln(hub.SkillLvl)
                if profile.Faceitlvl >= hub.SkillLvl {
                    log.Infoln("Getting invite..")
                    link, err := getInvite(hub.HubID)
                    if err != nil || link == "" {
                        log.Error(err)
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
