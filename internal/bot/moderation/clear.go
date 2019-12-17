package moderation

import (
    "github.com/arturoguerra/d2arena/internal/router"
    "strings"
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "gopkg.in/go-playground/validator.v9"
)

func deleteProfile(id string, s int) (*Profile, error) {
    base := "https://destinyarena.fireteamsupport.net/infoexchange.php?key=2YHSbPt5GJ9Uupgk"

    switch s {
        case 0:
            // Discord
            base += "&d=true&discordid=" + id
        case 1:
            // Faceit ID
            base += "&f=true&faceitguid=" + id
        case 2:
            // Faceit name
            base += "&f=true&faceitname=" + id
    }

    base += "&r=true"

    fmt.Println(base)
    req, _ := http.NewRequest("GET", base, nil)
    resp, err := requests.Internal.Do(req)

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


func Delete(ctx *router.Context) {
    if ctx.ChannelID != StaffChannelID {
        return
    }

    var uid string

    if len(ctx.Mentions) > 0 {
        uid = ctx.Mentions[0].ID
    } else {
        split := strings.Split(ctx.Content, " ")
        uid = split[1]
    }

    fmt.Println(ctx.Content)
    fmt.Println(uid)

    if uid == "" {
        return
    }

    idtype := sortProfileId(uid)

    profile, err := deleteProfile(uid, idtype)
    if err != nil {
        ctx.Session.ChannelMessageSend(ctx.ChannelID, "Error Deleting Profile")
        return
    }

    ctx.Session.ChannelMessageSend(ctx.ChannelID, "Deleted: \n" + generateProfile(ctx, profile))
}
