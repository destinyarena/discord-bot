package router

import "github.com/bwmarrin/discordgo"

type (
	ComponentContext struct {
		Session     *discordgo.Session
		Interaction *discordgo.Interaction
		Message     *discordgo.Message
	}

	ComponentHandlerFunc func(ctx *ComponentContext)

	Component struct {
		Name    string
		Type    discordgo.ComponentType
		Handler ComponentHandlerFunc
	}
)
