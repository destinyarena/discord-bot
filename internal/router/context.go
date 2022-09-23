package router

import "github.com/bwmarrin/discordgo"

type (
	Context struct {
		Session           *discordgo.Session
		Interaction       *discordgo.Interaction
		InteractionCreate *discordgo.InteractionCreate
		Options           map[string]*discordgo.ApplicationCommandInteractionDataOption
	}
)
