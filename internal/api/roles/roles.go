package roles

import (
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
    "github.com/arturoguerra/d2arena/internal/structs"
    "github.com/arturoguerra/d2arena/internal/config"
    "net/http"
    "strings"
    "gopkg.in/go-playground/validator.v9"
    "fmt"
)


type Hub struct {
    Id string `validate:"required"`
    Format string `validate:"required"`
}

var hubs [2]Hub

func init() {
    faceit := config.LoadFaceit()

    general := Hub{
        faceit.GeneralDiv,
        "Destiny Arena Division: {invite}\n",
    }

    doubles := Hub{
        faceit.DoublesDiv,
        "Destiny Doubles Division: {invite}\n",
    }

    hubs = [2]Hub{general, doubles}
}

func checkHub(hubid string, guid string) bool {
    return false
}


func sendInvites(s *discordgo.Session, guildid string, p *structs.RolesPayload, cfg *structs.Discord) {
    v := validator.New()
    message := "Invite links to join FACEIT Hubs.\n\n"
    send := false
    for _, hub := range hubs {
        if err := v.Struct(hub); err == nil {
            if inhub := checkHub(hub.Id, p.Faceit); inhub == false {
                if link, _ := sendLink(hub.Id); link != "" {
                    message += strings.Replace(hub.Format, "{invite}", link, 1)
                    send = true
                }
            }
        }
    }

    if send {
        s.ChannelMessageSend(p.Discord, message)
    }

    s.GuildMemberRoleAdd(cfg.GuildID, p.Discord, cfg.FaceitRoleID)
}

func New(s *discordgo.Session) echo.HandlerFunc {
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

        if err = v.Struct(payload); err != nil {
            return c.String(http.StatusBadRequest, "Error invalid payload")
        }

        fmt.Println(payload)


        sendInvites(s, g.ID, payload, discord)
        return c.String(http.StatusOK, "Roles have been assigned && and invites have been sent")
    }
}
