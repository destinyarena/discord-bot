package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/bot/internal/background"
	"github.com/destinyarena/bot/internal/bot/commands"
	"github.com/destinyarena/bot/internal/bot/handlers"
	"github.com/destinyarena/bot/internal/config"
	"github.com/destinyarena/bot/internal/logging"
	"github.com/destinyarena/bot/internal/natsevents"
	"github.com/destinyarena/bot/internal/router"
	"github.com/destinyarena/bot/internal/structs"
	"github.com/nats-io/nats.go"
)

const discordRegistration = "registration"

func main() {
	log := logging.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error(err)
		return
	}

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

	dgo.AddHandler(r.EventHandler)
	dgo.AddHandler(h.OnReady)
	dgo.AddHandler(h.OnMessageReactionAdd)
	dgo.AddHandler(h.OnMessageReactionRemove)
	dgo.AddHandler(h.OnMemberJoin)

	b := background.New(dgo, cfg)

	go b.UpdateRoles()

	// jonpall#2554 this is for you :)

	dgo.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	err = dgo.Open()
	if err != nil {
		log.Error(err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dgo.Close()
}
