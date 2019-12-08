package api

import (
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/arturoguerra/d2arena/internal/structs"
    "net/http"
    "bytes"
    "encoding/json"
    "io/ioutil"
    "fmt"
    "strconv"
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

func sendLink(s *discordgo.Session, hub int, uid string) {
    link, err := GenerateLink(hub)
    if err != nil {
        return
    }

    channel, err := s.UserChannelCreate(uid)
    if err != nil {
        return
    }

    s.ChannelMessageSend(channel.ID, link)
}


func New(e *echo.Echo, s *discordgo.Session) {
    e.POST("/roles", rolesFunc(s));
}


func phpHotFix(id string) int {
    n, _ := strconv.Atoi(id)
    return n
}

func updateRoles(s *discordgo.Session, gid string, p *structs.RolesPayload, cfg *structs.Discord) {
    lvl := phpHotFix(p.Skillvl)
    if lvl >= 8 {
        sendLink(s, hubs.DIV1, p.Discord)
        s.GuildMemberRoleAdd(gid, p.Discord, cfg.Div1)

    }

    if lvl >= 4 {
        sendLink(s, hubs.DIV2, p.Discord)
        s.GuildMemberRoleAdd(gid, p.Discord, cfg.Div2)
    }

    if 3 >= lvl {
        sendLink(s, hubs.DIV3, p.Discord)
        s.GuildMemberRoleAdd(gid, p.Discord, cfg.Div3)
    }
}

func rolesFunc(s *discordgo.Session) echo.HandlerFunc {
    discord := config.LoadDiscord()
    authtoken := config.LoadAuth()
    return func(c echo.Context) error {
        auth := c.Request().Header.Get("authentication")
        if auth != "Basic " + authtoken {
            return c.String(401, "Invalid Token")
        }

        g, err := s.Guild(discord.GuildID)
        if err != nil {
            fmt.Println(err)
            return c.String(500, "Well shit we fucked up")
        }

        payload := new(structs.RolesPayload)

        if err := c.Bind(payload); err != nil {
            return c.String(http.StatusBadRequest, "Error invalid payload")
        }

        updateRoles(s, g.ID, payload, discord)
        return c.String(http.StatusOK, "Roles have been assigned")
    }
}
