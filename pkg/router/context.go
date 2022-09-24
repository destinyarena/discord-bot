package router

import "github.com/bwmarrin/discordgo"

type (
	Context struct {
		Session     *discordgo.Session
		Interaction *discordgo.Interaction
		Options     map[string]*discordgo.ApplicationCommandInteractionDataOption
	}
)

func (c *Context) Reply(content string) error {
	err := c.Session.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})

	return err
}

func (c *Context) ReplyEmbed(embed *discordgo.MessageEmbed) error {
	err := c.Session.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})

	return err
}

func (c *Context) ReplyEmbeds(embeds []*discordgo.MessageEmbed) error {
	err := c.Session.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embeds,
		},
	})

	return err
}
