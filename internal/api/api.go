package api

import (
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
    "github.com/arturoguerra/d2arena/internal/api/roles"
)


func New(e *echo.Echo, s *discordgo.Session) {
    e.POST("/roles", roles.New(s));
}
