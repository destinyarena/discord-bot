package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/arturoguerra/faceitgo"
	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/discord-bot/internal/commands"
	"github.com/destinyarena/discord-bot/pkg/router"
)

func main() {
	session, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting router")

	router, err := router.NewRouter()
	if err != nil {
		panic(err)
	}

	faceit := faceitgo.New(&faceitgo.RESTConfig{
		Token:     os.Getenv("FACEIT_TOKEN"),
		PrivToken: os.Getenv("FACEIT_PRIV_TOKEN"),
	})

	fmt.Println("Adding commands..")

	if _, err = commands.New(router, faceit); err != nil {
		panic(err)
	}

	session.AddHandler(router.Handler)

	err = session.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println("Bot is running")

	if err := router.Sync(session); err != nil {
		panic(err)
	}

	defer session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("Started")
	<-stop
}
