package utils

import "github.com/bwmarrin/discordgo"

func NewButton(customID, label string, style discordgo.ButtonStyle) *discordgo.Button {
	return &discordgo.Button{
		Label:    label,
		Style:    style,
		CustomID: customID,
	}
}

func NewButtonWithEmoji(customID, label string, style discordgo.ButtonStyle, emoji discordgo.ComponentEmoji) *discordgo.Button {
	return &discordgo.Button{
		Label:    label,
		Style:    style,
		CustomID: customID,
		Emoji:    emoji,
	}
}
