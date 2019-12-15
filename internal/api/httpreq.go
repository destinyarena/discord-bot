package api

import (
    "net/http"
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/arturoguerra/d2arena/internal/structs"
    "encoding/json"
    "io/ioutil"
    "bytes"
    "github.com/bwmarrin/discordgo"
    "fmt"
    "errors"
)

func genLink(hubid string) (string, error) {
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

    if body.Payload.Code == "" {
        err = errors.New("Invalid invite code")
        return "", err
    }

    link := "https://www.faceit.com/en/inv/" + body.Payload.Code

    return link, nil
}

func sendLink(s *discordgo.Session, hubid string, uid string) error {
    link, err := genLink(hubid)
    if err != nil {
        return err
    }

    channel, err := s.UserChannelCreate(uid)
    if err != nil {
        return err
    }

    if _, err := s.ChannelMessageSend(channel.ID, link); err != nil {
        fmt.Println("Error sending invite to :" + uid)
        return err
    } else {
        fmt.Println("Sending " + link + " to " + uid)
        return nil
    }
}
