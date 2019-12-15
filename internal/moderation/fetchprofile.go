package moderation

import (
    "net/http"
    "encoding/json"
    "io/ioutil"
    "fmt"
    "gopkg.in/go-playground/validator.v9"
)

type Profile struct {
    DiscordID string `json:"discordid" validate:"required"`
    SteamID string `json:"steamid" validate:"required"`
    FaceitGuid string `json:"faceitguid" validate:"required"`
    FaceitName string `json:"faceitname" valitate:"required"`
}

func fetchProfile(id string) (*Profile, error) {
    base := "https://destinyarena.fireteamsupport.net/infoexchange.php?key=2YHSbPt5GJ9Uupgk&d=true&discordid=" + id
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
