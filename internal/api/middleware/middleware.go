package mdlware

import (
    "fmt"
    "github.com/labstack/echo"
    "github.com/arturoguerra/d2arena/internal/config"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
    token := config.LoadAuth()
    return func(c echo.Context) error {
        if token == "" {
            return c.String(500, "Missing Token")
        }

        auth := c.Request().Header.Get("Authorization")
        if auth != "Bearer " + token {
            return c.String(401, "Unauthorized access")
        } else {
            fmt.Println("Front-end Authenticated with API")
            return next(c)
        }
    }
}
