package bans

import (
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
    "net/http"
)

func New(s *discordgo.Session) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.String(http.StatusOK, "Herro")
    }
}
