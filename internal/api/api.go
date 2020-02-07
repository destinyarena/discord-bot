package api

import (
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
    "github.com/arturoguerra/d2arena/internal/api/middleware"
    "github.com/arturoguerra/d2arena/internal/api/roles"
    "github.com/arturoguerra/d2arena/internal/api/users"
)


func New(e *echo.Echo, s *discordgo.Session) {
    e.POST("/roles/:id", roles.New(s), mdlware.Auth);

    e.GET("/users/:id", users.New(s), mdlware.Auth);
}
