package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/discord-bot/internal/commands"
)

func main() {
	session, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	router, err := commands.New()
	if err != nil {
		panic(err)
	}

	session.AddHandler(router.Handler)

	err = session.Open()
	if err != nil {
		panic(err)
	}

	if err := router.Sync(session, ""); err != nil {
		panic(err)
	}

	defer session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("Started")
	<-stop
}
