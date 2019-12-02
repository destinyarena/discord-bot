package main

import (
    "github.com/arturoguerra/d2arena/internal/router"
    "github.com/arturoguerra/d2arena/internal/handlers"
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/arturoguerra/d2arena/internal/commands"
    "github.com/arturoguerra/d2arena/internal/background"
    "github.com/bwmarrin/discordgo"
    "fmt"
    "os"
    "os/signal"
    "syscall"
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

    commands.New(r)

    dgo.AddHandler(func (s *discordgo.Session, m *discordgo.MessageCreate) {
        r.Handler(m)
        //handlers.OnMessage(s, m)
    })

    dgo.AddHandler(handlers.OnReady)
    dgo.AddHandler(handlers.OnMemberJoin)

    go background.UpdateRoles(dgo)

    err = dgo.Open()
    if err != nil {
        fmt.Println(err)
        return
    }

    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <- sc

    dgo.Close()
}
