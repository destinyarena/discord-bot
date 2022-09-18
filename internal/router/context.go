package router

import "github.com/bwmarrin/discordgo"

type (
	Context struct {
		Session     *discordgo.Session
		Interaction *discordgo.InteractionCreate
	}
)
