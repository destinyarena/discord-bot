package bans

import (
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
    "net/http"
    "gopkg.in/go-playground/validator.v9"
)

type User struct {
    Id string `validate:"required"`
}

func New(s *discordgo.Session) echo.HandlerFunc {
    return func(c echo.Context) error {
        u := new(User)
        if err := c.Bind(u); err != nil {
            return c.String(http.StatusBadRequest, "Error invaild payload")
        }

        v := validator.New()
        if err := v.Struct(u); err != nil {
            return c.String(http.StatusBadRequest, "Error invalid payload")
        }


        return c.String(http.StatusOK, "User had been banned")
    }
}
