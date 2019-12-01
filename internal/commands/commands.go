package commands

import (
    "github.com/arturoguerra/d2arena/internal/router"
    "github.com/arturoguerra/d2arena/internal/structs"
    "bytes"
    "encoding/json"
    "net/http"
    "io/ioutil"
)

type HUBS struct {
    DIV1 int
    DIV2 int
    DIV3 int
}


var hubs *HUBS

func init() {
    hubs = &HUBS{
        0,
        1,
        2,
    }
}

func New(r *router.Router) {
    r.On("div1", getLink(hubs.DIV1))
    r.On("div2", getLink(hubs.DIV2))
    r.On("div3", getLink(hubs.DIV3))
}


func GenerateLink(hub int) (string, error) {
    var hubid string

    switch hub {
        case hubs.DIV1:
            hubid = "DIV1"
        case hubs.DIV2:
            hubid = "DIV2"
        case hubs.DIV3:
            hubid = "484d1f7e-ad44-417e-b1c8-4a10c1159808"
    }

    reqBody, _ := json.Marshal(structs.ReqBody{
        hubid,
        "hub",
        "regular",
        1800,
        1,
    })

    client := &http.Client{}

    req, _ := http.NewRequest("POST", "https://api.faceit.com/invitations/v1/invite", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer ") //+ faceitcfg.UserToken)

    resp, err := client.Do(req)
    defer resp.Body.Close()

    if err != nil {
        return "", err
    }

    link, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(link), nil
}

func getLink(hub int) func(ctx *router.Context) {
    return func (ctx *router.Context) {
        link, err := GenerateLink(hub)
        if err != nil {
            return
        }

        channel, err := ctx.Session.UserChannelCreate(ctx.Message.Author.ID)
        if err != nil {
            return
        }

        ctx.Session.ChannelMessageSend(channel.ID, link)
    }
}