package commands

import (
    "github.com/arturoguerra/d2arena/internal/router"
    "github.com/arturoguerra/d2arena/internal/structs"
    "github.com/arturoguerra/d2arena/internal/config"
    "bytes"
    "encoding/json"
    "net/http"
    "io/ioutil"
)

type HUBS struct {
    DIV1 int
    DIV2 int
    DIV3 int
    DUEL int
    STADIUM int
}


var hubs *HUBS

func init() {
    hubs = &HUBS{
        0,
        1,
        2,
        3,
        4,
    }
}

func New(r *router.Router) {
    //r.On("div1", getLink(hubs.DIV1))
    //r.On("div2", getLink(hubs.DIV2))
    //r.On("div3", getLink(hubs.DIV3))
    //r.On("public", getLink(hubs.DIV3))
    //r.On("divd", getLink(hubs.DUEL))
    //r.On("divs", getLink(hubs.STADIUM))
}


func GenerateLink(hub int) (string, error) {
    var hubid string

    switch hub {
        case hubs.DIV1:
            hubid = "0ced849e-e10b-4998-8780-d85c60135f9d"
        case hubs.DIV2:
            hubid = "cf70962d-756f-4c54-9492-7cc06b33d685"
        case hubs.DIV3:
            hubid = "e615de71-bea1-4d5b-9e0e-d14f410164d4"
        case hubs.DUEL:
            hubid = "ea3a5dbe-e85f-4ebe-9c56-062e1a3160f2"
        case hubs.STADIUM:
            hubid = "2133e0f1-523a-41ae-a41a-1686f0ba1528"
    }

    reqBody, _ := json.Marshal(structs.ReqBody{
        hubid,
        "hub",
        "regular",
        1800,
        1,
    })

    client := &http.Client{}

    fitcfg := config.LoadFaceit()

    req, _ := http.NewRequest("POST", "https://api.faceit.com/invitations/v1/invite", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer " + fitcfg.UserToken)

    resp, err := client.Do(req)
    defer resp.Body.Close()

    if err != nil {
        return "", err
    }

    rawbody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var body structs.ResponseBody
    json.Unmarshal([]byte(rawbody), &body)

    link := "https://www.faceit.com/en/inv/" + body.Payload.Code

    return link, nil
}

func getLink(hub int) func(ctx *router.Context) {
    return func (ctx *router.Context) {
        if ctx.Channel.ID != "650159281575821312" {
            return
        }

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
