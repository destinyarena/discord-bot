package api

import (
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/arturoguerra/d2arena/internal/structs"
    "net/http"
    "fmt"
)

func New(e *echo.Echo, s *discordgo.Session) {
    e.POST("/roles", rolesFunc(s));
}

func updateRoles(s *discordgo.Session, gid string, p *structs.RolesPayload, cfg *structs.Discord) {
    if p.Skillvl >= 8 {
        s.GuildMemberRoleAdd(gid, p.Discord, cfg.Div1)
    }

    if p.Skillvl >= 4 {
        s.GuildMemberRoleAdd(gid, p.Discord, cfg.Div2)
    }

    if 3 >= p.Skillvl {
        s.GuildMemberRoleAdd(gid, p.Discord, cfg.Div3)
    }
}

func test(s string) {
    fmt.Println(s);
}

func rolesFunc(s *discordgo.Session) echo.HandlerFunc {
    discord := config.LoadDiscord()
    return func(c echo.Context) error {
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
