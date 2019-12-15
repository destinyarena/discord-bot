package api

import (
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
    "github.com/arturoguerra/d2arena/internal/structs"
    "github.com/arturoguerra/d2arena/internal/config"
    "net/http"
    "gopkg.in/go-playground/validator.v9"
    "fmt"
)


type Hub struct {
    Id string
    RoleID string
}

var hubs = [2]Hub{}

func init() {
    general := Hub{
    }

    stadium := Hub{
    }

    hubs = [2]Hub{general, stadium}
}

func checkHub(hubid string) bool {
    return false
}


func updateRoles(s *discordgo.Session, gid string, p *structs.RolesPayload, cfg *structs.Discord) {
        for _, hub := range hubs {
            if inhub := checkHub(p.Faceit); inhub == false {
                sendLink(s, hub.Id, p.Discord)
            }

            s.GuildMemberRoleAdd(gid, p.Discord, hub.RoleID)
        }
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

        err = v.Struct(payload)
        fmt.Println(err)

        fmt.Println(payload)


        updateRoles(s, g.ID, payload, discord)
        return c.String(http.StatusOK, "Roles have been assigned")
    }
}
