package api

import (
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/arturoguerra/d2arena/internal/structs"
    "net/http"
)

func New(e *echo.Echo, s *discordgo.Session) {
    e.POST("/roles", rolesFunc(s));
}

func updateRoles(s *discordgo.Session, gid string, p *structs.RolesPayload, cfg *structs.Discord) error {
    var err error;
    if p.Skillvl >= 8 {
        err = s.GuildMemberRoleAdd(gid, p.Discord, cfg.Div1);
    }

    if p.Skillvl >= 4 {
        err = s.GuildMemberRoleAdd(gid, p.Discord, cfg.Div2);
    }

    if 3 >= p.Skillvl {
        err = s.GuildMemberRoleAdd(gid, p.Discord, cfg.Div3);
    }

    return err;
}

func rolesFunc(s *discordgo.Session) echo.HandlerFunc {
    discord := config.LoadDiscord();
    return func(c echo.Context) error {
        g, _ := s.Guild(discord.GuildID);
        payload := new(structs.RolesPayload);

        if err := c.Bind(payload); err != nil {
            return c.String(http.StatusBadRequest, "Error invalid payload");
        }

        if err := updateRoles(s, g.ID, payload, discord); err != nil {
            return c.String(http.StatusUnauthorized, "Error unauthorized payload");
        } else {
            return c.String(http.StatusOK, "Roles have been assigned");
        }
    }
}
