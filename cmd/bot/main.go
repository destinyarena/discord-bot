package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arturoguerra/d2arena/internal/background"
	"github.com/arturoguerra/d2arena/internal/bot/commands"
	"github.com/arturoguerra/d2arena/internal/bot/handlers"
	"github.com/arturoguerra/d2arena/internal/config"
	"github.com/arturoguerra/d2arena/internal/logging"
	"github.com/arturoguerra/d2arena/internal/natsevents"
	"github.com/arturoguerra/d2arena/internal/router"
	"github.com/arturoguerra/d2arena/internal/structs"
	"github.com/bwmarrin/discordgo"
	"github.com/labstack/echo"
	"github.com/nats-io/nats.go"
)

const discordRegistration = "registration"

func main() {
	log := logging.New()

	cfg, err := config.LoadConfig()

	log.Infof("Starting NATS Client: %s", cfg.NATS.URL)
	nc, err := nats.Connect(cfg.NATS.URL)
	if err != nil {
		log.Fatal(err)
	}

	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	recvRegistrationChan := make(chan *structs.NATSRegistration)
	ec.BindRecvChan(discordRegistration, recvRegistrationChan)

	nchan := &structs.NATS{
		RecvRegistration: recvRegistrationChan,
	}

	log.Infoln("Starting Discord Bot")

	dgo, err := discordgo.New("Bot " + cfg.Discord.Token)
	if err != nil {
		log.Error(err)
		return
	}

	r := router.New(cfg, log)

	commands.New(r)

	h := handlers.New(cfg, log)

	natsevents.New(dgo, cfg, log, nchan)

	// Command Handler
	dgo.AddHandler(r.EventHandler)

	dgo.AddHandler(h.OnReady)
	dgo.AddHandler(h.OnMemberJoin)
	dgo.AddHandler(h.OnMessageReactionAdd)
	dgo.AddHandler(h.OnMessageReactionRemove)

	b := background.New(dgo, cfg)

	go b.UpdateRoles()

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
	<-sc

	dgo.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
