package users

import (
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
    "net/http"
)

type ReturnUser struct {
    ID       string   `json:"id"`
    Username string   `json:"username"`
    Roles    []string `json:"roles"`
}

func New(s *discordgo.Session) echo.HandlerFunc {
    return func(c echo.Context) error {
        id := c.Param("id")

        cfg := config.LoadDiscord()

        member, err := s.GuildMember(cfg.GuildID, id)
        if err != nil {
            return err
        }

        p := &ReturnUser{
            ID:       member.User.ID,
            Username: member.User.Username,
            Roles:    member.Roles,
        }

        return c.JSON(http.StatusOK, p)
    }
}
