package router

import "github.com/bwmarrin/discordgo"

type (
	CommandInterface interface {
		GetName() string
		GetDescription() string
		GetApplicationCommand() *discordgo.ApplicationCommand
		GetApplicationSubCommands() []*discordgo.ApplicationCommandOption
		GetOptions() []*discordgo.ApplicationCommandOption
		GetSubCommands() []CommandInterface
		GetSubCommand(name string) (CommandInterface, bool)
		Handler(ctx *Context)
	}

	BasicCommand struct {
		Name        string
		Description string
		SubCommands []CommandInterface
		Options     []*discordgo.ApplicationCommandOption
	}
)

func (c *BasicCommand) GetName() string {
	return c.Name
}

func (c *BasicCommand) GetDescription() string {
	return c.Description
}

func (c *BasicCommand) GetOptions() []*discordgo.ApplicationCommandOption {
	return c.Options
}

func (c *BasicCommand) GetSubCommand(name string) (CommandInterface, bool) {
	for _, sub := range c.SubCommands {
		if sub.GetName() == name {
			return sub, true
		}
	}

	return nil, false
}

func (c *BasicCommand) GetSubCommands() []CommandInterface {
	return c.SubCommands
}

func (c *BasicCommand) GetApplicationCommand() *discordgo.ApplicationCommand {
	ac := &discordgo.ApplicationCommand{
		Name:        c.GetName(),
		Description: c.GetDescription(),
		Type:        discordgo.ChatApplicationCommand,
		Options:     c.Options,
	}

	ac.Options = append(ac.Options, c.GetApplicationSubCommands()...)

	return ac
}

func (c *BasicCommand) GetApplicationSubCommands() []*discordgo.ApplicationCommandOption {
	options := make([]*discordgo.ApplicationCommandOption, 0)

	for _, subcommand := range c.SubCommands {
		if len(subcommand.GetSubCommands()) != 0 {
			options = append(options, &discordgo.ApplicationCommandOption{
				Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
				Name:        subcommand.GetName(),
				Description: subcommand.GetDescription(),
				Options:     subcommand.GetApplicationSubCommands(),
			})
		} else {
			options = append(options, &discordgo.ApplicationCommandOption{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        subcommand.GetName(),
				Description: subcommand.GetDescription(),
				Options:     subcommand.GetOptions(),
			})
		}
	}

	return options
}
