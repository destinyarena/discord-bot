package main

import (
	"github.com/andersfylling/disgord"
	"github.com/destinyarena/bot/internal/bot"
	"github.com/destinyarena/bot/internal/commands"
	"github.com/destinyarena/bot/internal/config"
	"github.com/destinyarena/bot/internal/faceit"
	"github.com/destinyarena/bot/internal/handlers"
	"github.com/destinyarena/bot/internal/logging"
	"github.com/destinyarena/bot/internal/natsevents"
	"github.com/destinyarena/bot/internal/profiles"
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

	dclient := disgord.New(disgord.Config{
		BotToken: cfg.Discord.Token,
		Logger:   log,
	})

	defer dclient.Gateway().StayConnectedUntilInterrupted()

	log.Info(cfg.Discord)
	log.Info(cfg.Profiles)

	bot, err := bot.New(log, dclient, cfg.Discord)
	if err != nil {
		log.Fatal(err)
	}

	fclient, err := faceit.New(log, cfg.Faceit)
	if err != nil {
		log.Fatal(err)
	}

	pclient, err := profiles.New(log, cfg.Profiles)
	if err != nil {
		log.Fatal(err)
	}

	// jonpall#2554 this is for you :)
	commands.New(bot, fclient, pclient, log)
	handlers.New(bot, fclient, pclient, log)
	natsevents.New(bot, log, nchan)

}
