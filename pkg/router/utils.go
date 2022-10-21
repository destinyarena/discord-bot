package router

import (
	"github.com/bwmarrin/discordgo"
)

func convertOptionsToMap(options []*discordgo.ApplicationCommandInteractionDataOption) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	m := make(map[string]*discordgo.ApplicationCommandInteractionDataOption)
	for _, o := range options {
		m[o.Name] = o
	}

	return m
}
