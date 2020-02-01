package handlers

import (
    "fmt"
    "net/http"
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/arturoguerra/d2arena/internal/structs"
    "encoding/json"
    "io/ioutil"
    "bytes"
    "errors"
)

func getInvite(hubid string) (string, error) {
    fmt.Println(hubid)
    reqBody, _ := json.Marshal(structs.ReqBody{
        hubid,
        "hub",
        "regular",
        0,
        1,
    })

    client := &http.Client{}

    fitcfg := config.LoadFaceit()
    fmt.Println(fitcfg.UserToken)
    req, _ := http.NewRequest("POST", "https://api.faceit.com/invitations/v1/invite", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer " + fitcfg.UserToken)
    resp, err := client.Do(req)
    defer resp.Body.Close()

    if err != nil {
        return "", err
    }

    if resp.StatusCode != 200 && resp.StatusCode != 201 {
        err = fmt.Errorf("Server response code: %d", resp.StatusCode)
        return "", err
    }

    rawbody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    fmt.Sprintf(string(rawbody))

    var body structs.ResponseBody
    json.Unmarshal([]byte(rawbody), &body)

    if body.Payload.Code == "" {
        err = errors.New("Invalid invite code")
        return "", err
    }

    link := "https://www.faceit.com/en/inv/" + body.Payload.Code

    return link, nil
}
