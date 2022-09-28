package router

import "github.com/bwmarrin/discordgo"

type (
	ContextInterface interface {
		Reply(content string, embeds []*discordgo.MessageEmbed, components []discordgo.MessageComponent) error
		UpdateReply(content string, embeds []*discordgo.MessageEmbed, components []discordgo.MessageComponent) error
		EphemeralReply(content string, embeds []*discordgo.MessageEmbed, components []discordgo.MessageComponent) error
		GetSession() *discordgo.Session
		GetInteraction() *discordgo.Interaction
		GetMessage() *discordgo.Message
		GetOptions() map[string]*discordgo.ApplicationCommandInteractionDataOption
	}

	Context struct {
		Session     *discordgo.Session
		Interaction *discordgo.Interaction
		Message     *discordgo.Message
		Options     map[string]*discordgo.ApplicationCommandInteractionDataOption
	}
)

func NewContext(s *discordgo.Session, i *discordgo.Interaction) *Context {
	return &Context{
		Session:     s,
		Interaction: i,
		Message:     i.Message,
		Options:     make(map[string]*discordgo.ApplicationCommandInteractionDataOption),
	}
}

func (c *Context) GetSession() *discordgo.Session {
	return c.Session
}

func (c *Context) GetInteraction() *discordgo.Interaction {
	return c.Interaction
}

func (c *Context) GetMessage() *discordgo.Message {
	return c.Message
}

func (c *Context) GetOptions() map[string]*discordgo.ApplicationCommandInteractionDataOption {
	return c.Options
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
