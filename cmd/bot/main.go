package main

import (
    //"github.com/arturoguerra/d2arena/internal/structs"
    "github.com/arturoguerra/d2arena/internal/router"
    "github.com/arturoguerra/d2arena/internal/handlers"
    //"github.com/arturoguerra/d2arena/internal/config"
    "github.com/arturoguerra/d2arena/internal/commands"
    "github.com/bwmarrin/discordgo"
    "fmt"
    "os"
    "os/signal"
    "syscall"
)


func init() {
}

func main() {
    cfg := router.NewConfig(
        "",
        "",
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

    dgo.AddHandler(func (_ *discordgo.Session, m *discordgo.MessageCreate) {
        r.Handler(m)
    })

    dgo.AddHandler(handlers.OnReady)

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
