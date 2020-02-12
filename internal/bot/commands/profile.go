package commands

import (
    "fmt"
    "context"
    "google.golang.org/grpc"
    "github.com/bwmarrin/discordgo"
    "github.com/arturoguerra/d2arena/pkg/router"
    "github.com/arturoguerra/d2arena/pkg/faceit"
    "github.com/arturoguerra/d2arena/pkg/profiles"
    "github.com/arturoguerra/d2arena/internal/bot/utils"
)

type FaceitProfile struct {
    ID       string
    Username string
    Level    int
}

type Profile struct {
    Discord string
    Faceit  string
    Bungie  string
    Banned  bool
}

func getFaceitProfile(id string) (*FaceitProfile, error) {
    conn, err := grpc.Dial(grpcfaceit, grpc.WithInsecure())
    if err != nil {
        return nil, err
    }

    defer conn.Close()

    c := faceit.NewFaceitClient(conn)
    log.Info("Fetching faceit profile")
    r, err := c.GetProfile(context.Background(), &faceit.ProfileRequest{
        Guid: id,
    })
    if err != nil {
        return nil, err
    }

    return &FaceitProfile{
        Username: r.GetUsername(),
        ID:       r.GetGuid(),
        Level:    int(r.GetSkilllvl()),
    }, nil
}

func getProfile(id string) (*Profile, error) {
    conn, err := grpc.Dial(grpcprofiles, grpc.WithInsecure())
    if err != nil {
        return nil, err
    }

    defer conn.Close()

    c := profiles.NewProfilesClient(conn)
    log.Info("Fetching Profile from database")
    r, err := c.GetProfile(context.Background(), &profiles.IdRequest{
        Id: id,
    })
    if err != nil {
        return nil, err
    }

    return &Profile{
        Discord: r.GetDiscord(),
        Faceit:  r.GetFaceit(),
        Bungie:  r.GetBungie(),
        Banned:  r.GetBanned(),
    }, nil
}

func profile(ctx *router.Context) {
    var embed *discordgo.MessageEmbed
    var uid string


    err, uid := utils.GetUID(ctx)
    if err != nil {
        embed = &discordgo.MessageEmbed{
            Title: "Error fetching Profile",
            Description: err.Error(),
            Color: 0xf30707,
        }

        ctx.ReplyEmbed(embed)
        return
    }

    fmt.Println(uid)

    profile, err := getProfile(uid)
    if err != nil {
        embed = &discordgo.MessageEmbed{
            Title: "Error Fetching User",
            Description: "Looks like you provided an invaild id or the user never registered",
            Color: 0xf30707,
        }

        ctx.ReplyEmbed(embed)
        return
    }

    duser, err := ctx.Session.User(profile.Discord)
    if err != nil {
        log.Error(err)
        embed = &discordgo.MessageEmbed{
            Title: "Error Fetching User",
            Description: "Error fetching Profile",
            Color: 0xf30707,
        }
        ctx.ReplyEmbed(embed)
        return
    }

    fprofile, err := getFaceitProfile(profile.Faceit)
    if err != nil {
        log.Error(err)
        fprofile = &FaceitProfile{
            ID:       profile.Faceit,
            Username: "Unavailable",
            Level:    0,
        }
    }

    fields := make([]*discordgo.MessageEmbedField, 0)

    fields = append(fields, &discordgo.MessageEmbedField{
        Name: "Discord Username",
        Value: fmt.Sprintf("%s#%s", duser.Username, duser.Discriminator),
    })

    fields = append(fields, &discordgo.MessageEmbedField{
        Name: "Discord ID",
        Value: duser.ID,
    })

    fields = append(fields, &discordgo.MessageEmbedField{
        Name: "Bungie ID",
        Value: profile.Bungie,
    })

    fields = append(fields, &discordgo.MessageEmbedField{
        Name: "Faceit Username",
        Value: fprofile.Username,
    })

    fields = append(fields, &discordgo.MessageEmbedField{
        Name: "Faceit GUID",
        Value: fprofile.ID,
    })

    fields = append(fields, &discordgo.MessageEmbedField{
        Name: "Faceit Skill Level",
        Value: fmt.Sprintf("%d", fprofile.Level),
    })

    embed = &discordgo.MessageEmbed{
        Title: "User Profile",
        Color: 0x019fd8,
        Fields: fields,
    }

    ctx.ReplyEmbed(embed)
}
