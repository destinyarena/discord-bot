package main

import (
    "github.com/nats-io/nats.go"
    "github.com/arturoguerra/d2arena/internal/router"
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/arturoguerra/d2arena/internal/structs"
    "github.com/arturoguerra/d2arena/internal/background"
    "github.com/arturoguerra/d2arena/internal/logging"
    "github.com/arturoguerra/d2arena/internal/bot/handlers"
    "github.com/arturoguerra/d2arena/internal/bot/commands"
    "github.com/arturoguerra/d2arena/internal/natsevents"
    "github.com/bwmarrin/discordgo"
    "os"
    "github.com/labstack/echo"
    "os/signal"
    "net/http"
    "syscall"
    "context"
    "time"
)

const DISCORD_REGISTRATION = "registration"

func main() {
    log := logging.New()

    ncfg := config.LoadNATSConfig()
    log.Infof("Starting NATS Client: %s", ncfg.URL)
    nc, err := nats.Connect(ncfg.URL)
    if err != nil {
        log.Fatal(err)
    }

    ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

    recvRegistrationChan := make(chan *structs.NATSRegistration)
    ec.BindRecvChan(DISCORD_REGISTRATION, recvRegistrationChan)

    nchan := &structs.NATS{
        RecvRegistration: recvRegistrationChan,
    }

    log.Infoln("Starting Discord Bot")
    dcfg := config.LoadDiscord()
    dgo, err := discordgo.New("Bot " + dcfg.Token)

    if err != nil {
        log.Error(err)
        return
    }


    r:= router.New()

    commands.New(r)

    natsevents.New(dgo, nchan)

    // Command Handler
    dgo.AddHandler(r.EventHandler)

    dgo.AddHandler(handlers.OnReady)
    dgo.AddHandler(handlers.OnMemberJoin)
    dgo.AddHandler(handlers.OnMessageReactionAdd)
    dgo.AddHandler(handlers.OnMessageReactionRemove)

    go background.UpdateRoles(dgo)

    err = dgo.Open()
    if err != nil {
        log.Error(err)
        return
    }

    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "all good")
    })


    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    go func() {
        if err := e.Start(":" + port); err != nil {
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
