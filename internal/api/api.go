package api

import (
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
)


func New(e *echo.Echo, s *discordgo.Session) {
    e.POST("/roles", rolesFunc(s));
}
