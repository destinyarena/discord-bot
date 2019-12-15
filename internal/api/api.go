package api

import (
    "gopkg.in/go-playground/validator.v9"
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/arturoguerra/d2arena/internal/structs"
    "net/http"
    "bytes"
    "encoding/json"
    "io/ioutil"
    "fmt"
)

func GenerateLink(hubid string) (string, error) {
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

func sendLink(s *discordgo.Session, hubid string, uid string) {
    link, err := GenerateLink(hubid)
    if err != nil {
        return
    }

    channel, err := s.UserChannelCreate(uid)
    if err != nil {
        return
    }

    if _, err := s.ChannelMessageSend(channel.ID, link); err != nil {
        fmt.Println("Error sending invite to :" + uid)
    } else {
        fmt.Println("Sending " + link + " to " + uid)
    }
}


func New(e *echo.Echo, s *discordgo.Session) {
    e.POST("/roles", rolesFunc(s));
}


func updateRoles(s *discordgo.Session, gid string, p *structs.RolesPayload, cfg *structs.Discord) {
        //sendLink(s, hubs.DIV1, p.Discord)
        //s.GuildMemberRoleAdd(gid, p.Discord, cfg.Div1)
}

func rolesFunc(s *discordgo.Session) echo.HandlerFunc {
    discord := config.LoadDiscord()
    authtoken := config.LoadAuth()
    return func(c echo.Context) error {
        if authtoken == "" {
            return c.String(500, "Missing token")
        }

        auth := c.Request().Header.Get("Authorization")
        if auth != "Basic " + authtoken {
            return c.String(401, "Invalid Token")
        }

        fmt.Println("Front-end authenticated")

        g, err := s.Guild(discord.GuildID)
        if err != nil {
            fmt.Println(err)
            return c.String(500, "Well shit we fucked up")
        }

        payload := new(structs.RolesPayload)
        if err := c.Bind(payload); err != nil {
            return c.String(http.StatusBadRequest, "Error invalid payload")
        }

        v := validator.New()

        if err := v.Struct(payload); err != nil {
            return c.String(http.StatusBadRequest, "Error invalid payload")
        }

        fmt.Println(payload)


        updateRoles(s, g.ID, payload, discord)
        return c.String(http.StatusOK, "Roles have been assigned")
    }
}
