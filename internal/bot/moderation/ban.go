package moderation

import (
    "github.com/arturoguerra/d2arena/internal/router"
    "github.com/arturoguerra/d2arena/internal/config"
    "strings"
    "fmt"
    "errors"
    "encoding/json"
    "bytes"
//    "io/ioutil"
    "net/http"
)

type BanPayload struct {
    HubId string `json:"hubId"`
    UserId string `json:"userId"`
    Reason string `json:"reason"`
}


func Ban(ctx *router.Context) {
    if len(ctx.Mentions) != 1 {
        return
    }

    if ctx.ChannelID != StaffChannelID {
        return
    }

    uid := ctx.Mentions[0].ID

    split := strings.Split(ctx.Content, " ")
    reason := strings.Join(split[2:], " ")

    profile, err := fetchProfile(uid, 0)
    if err != nil {
        ctx.Session.ChannelMessageSend(ctx.ChannelID, "Error fetching user profile")
        return
    }

    if errs := FaceitBan(profile.FaceitGuid, reason); len(errs) != 0 {
        ctx.Session.ChannelMessageSend(ctx.ChannelID, "Error banning user from faceit please do it manually!")
    }

    if err = DiscordBan(ctx, uid, reason); err != nil {
        ctx.Session.ChannelMessageSend(ctx.ChannelID, "Error banning user from discord please do it manually!")
    }

    if err == nil {
        ctx.Session.ChannelMessageSend(ctx.ChannelID, "Banned user from discord and faceit")
    }
}

func DiscordBan (ctx *router.Context, uid, reason string) error {
    err := ctx.Session.GuildBanCreate(ctx.GuildID, uid, 7)
    return err
}

func FaceitBan (id string, reason string) []error {
    faceit := config.LoadFaceit()
    gen := faceit.GeneralDiv
    doubles := faceit.DoublesDiv
    hubs := [2]string{gen, doubles}
    errs := []error{}
    for _, hub := range hubs {
        url := "https://api.faceit.com/hubs/v1/hub/" + hub + "/ban/" + id
        payload := BanPayload{
            hub,
            id,
            reason,
        }

        body, _ := json.Marshal(payload)
        req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
        req.Header.Add("Authorization", "Bearer " + faceit.UserToken)
        req.Header.Add("Content-Type", "application/json")
        resp, err := requests.Client.Do(req)
        defer resp.Body.Close()

        if err != nil || resp.StatusCode != 200 {
            fmt.Println("Error banning user from faceit")
            if err == nil {
                err = errors.New("Can't kick from faceit :(")
            }

            errs = append(errs, err)
        } else {
           fmt.Println("Banned user from faceit")
        }
    }

    return errs
}
