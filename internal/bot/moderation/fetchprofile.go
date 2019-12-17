package moderation

import (
    "net/http"
    "encoding/json"
    "io/ioutil"
    "fmt"
    "regexp"
    "gopkg.in/go-playground/validator.v9"
)

type Profile struct {
    DiscordID string `json:"discordid" validate:"required"`
    SteamID string `json:"steamid" validate:"required"`
    FaceitGuid string `json:"faceitguid" validate:"required"`
    FaceitName string `json:"faceitname" valitate:"required"`
}

func sortProfileId(id string) int {
    var returntype int

    if m, _ := regexp.Match(`\d+`, []byte(id)); m {
        returntype = 0
    } else if m, _ := regexp.Match(`([A-f0-9\-])+`, []byte(id)); m {
        returntype = 1
    } else {
        returntype = 2
    }

    return returntype
}

func fetchProfile(id string, s int) (*Profile, error) {
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
