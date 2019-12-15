package moderation

import (
    "github.com/arturoguerra/d2arena/internal/router"
    "github.com/arturoguerra/d2arena/internal/config"
    "strings"
    "fmt"
    "encoding/json"
    "bytes"
    "io/ioutil"
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

    profile, err := fetchProfile(uid)
    if err != nil {
        ctx.Session.ChannelMessageSend(ctx.ChannelID, "Error fetching user profile")
        return
    }

    DiscordBan(ctx, uid, reason)
    FaceitBan(profile.FaceitGuid, reason)
    ctx.Session.ChannelMessageSend(ctx.ChannelID, "Banned user from discord and faceit")
}

func DiscordBan (ctx *router.Context, uid, reason string) {
    ctx.Session.GuildBanCreate(ctx.GuildID, uid, 7)
}

func FaceitBan (id string, reason string) {
    faceit := config.LoadFaceit()
    gen := faceit.GeneralDiv
    doubles := faceit.DoublesDiv
    hubs := [2]string{gen,doubles}
    for _, hub := range hubs {
        url := "https://api.faceit.com/hubs/v1/hub/" + hub + "/ban/" + id
        payload := BanPayload{
            hub,
            id,
            reason,
        }

        body, _ := json.Marshal(payload)
        req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
        resp, err := requests.Api.Do(req)
        defer resp.Body.Close()

        rawbody, _ := ioutil.ReadAll(resp.Body)
        fmt.Println(string(rawbody))

        if err != nil {
            fmt.Println("Error banning user from faceit")
            return
        }

        fmt.Println("Banned user from faceit")
    }
}
