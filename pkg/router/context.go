package router

import "github.com/bwmarrin/discordgo"

type (
	Context struct {
		Router      *Router
		Session     *discordgo.Session
		Interaction *discordgo.Interaction
		Message     *discordgo.Message
	}
)

func NewContext(s *discordgo.Session, i *discordgo.Interaction) *Context {
	return &Context{
		Session:     s,
		Interaction: i,
		Message:     i.Message,
	}
}

func (c *Context) Reply(content string, embeds []*discordgo.MessageEmbed, components []discordgo.MessageComponent) error {
	return c.reply(discordgo.InteractionResponseChannelMessageWithSource, 0, content, embeds, components)
}

func (c *Context) UpdateReply(content string, embeds []*discordgo.MessageEmbed, components []discordgo.MessageComponent) error {
	return c.reply(discordgo.InteractionResponseUpdateMessage, 0, content, embeds, components)
}

func (c *Context) EphemeralReply(content string, embeds []*discordgo.MessageEmbed, components []discordgo.MessageComponent) error {
	return c.reply(discordgo.InteractionResponseChannelMessageWithSource, discordgo.MessageFlagsEphemeral, content, embeds, components)
}

func (c *Context) reply(irtype discordgo.InteractionResponseType, flags discordgo.MessageFlags, content string, embeds []*discordgo.MessageEmbed, components []discordgo.MessageComponent) error {
	err := c.Session.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
		Type: irtype,
		Data: &discordgo.InteractionResponseData{
			Content:    content,
			Embeds:     embeds,
			Flags:      flags,
			Components: components,
		},
	})

	return err
}
