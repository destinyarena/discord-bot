package main

import (
    "github.com/arturoguerra/d2arena/internal/api"
    "github.com/arturoguerra/d2arena/internal/router"
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/arturoguerra/d2arena/internal/background"
    //"github.com/arturoguerra/d2arena/internal/bot/commands"
    "github.com/arturoguerra/d2arena/internal/bot/handlers"
    "github.com/arturoguerra/d2arena/internal/bot/moderation"
    "github.com/bwmarrin/discordgo"
    "fmt"
    "os"
    "github.com/labstack/echo"
    "os/signal"
    "syscall"
    "context"
    "time"
)

func lolkillme() {
    for true {
        fmt.Println("HELLO ITS ME")
    }
}

func main() {
    dcfg := config.LoadDiscord()
    cfg := router.NewConfig(
        dcfg.Prefix,
        dcfg.Token,
    )

    dgo, err := discordgo.New("Bot " + cfg.Token)

    if err != nil {
        fmt.Println(err)
        return
    }


    r:= router.New(
        dgo,
        cfg,
    )

    //commands.New(r)
    moderation.New(r)

    dgo.AddHandler(func (s *discordgo.Session, m *discordgo.MessageCreate) {
        r.Handler(m)
        handlers.OnMessage(s, m)
    })

    dgo.AddHandler(handlers.OnReady)
    dgo.AddHandler(handlers.OnMemberJoin)

    go background.UpdateRoles(dgo)

    err = dgo.Open()
    if err != nil {
        fmt.Println(err)
        return
    }

    e := echo.New()

    api.New(e, dgo)

    go func() {
        if err := e.Start(":9080"); err != nil {
            e.Logger.Info("Shutting down the server!")
        }
    }()


    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <- sc

    dgo.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := e.Shutdown(ctx); err != nil {
        e.Logger.Fatal(err)
    }
}
