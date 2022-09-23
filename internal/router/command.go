package router

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type (
	CommandInterface interface {
		GetName() string
		GetSubCommandGroups() []SubCommandGroupInterface
		GetSubCommandGroup(name string) (SubCommandGroupInterface, bool)
		GetSubCommands() []SubCommandInterface
		GetSubCommand(name string) (SubCommandInterface, bool)
		GetApplicationCommand() *discordgo.ApplicationCommand
		Handler(ctx *Context)
	}

	SubCommandGroupInterface interface {
		GetName() string
		GetSubCommands() []SubCommandInterface
		GetSubCommand(name string) (SubCommandInterface, bool)
		GetApplicationCommandOption() *discordgo.ApplicationCommandOption
	}

	SubCommandInterface interface {
		GetName() string
		GetApplicationCommandOption() *discordgo.ApplicationCommandOption
		Handler(ctx *Context)
	}

	BaseCommand struct {
		Name             string
		Description      string
		SubCommands      []SubCommandInterface
		SubCommandGroups []SubCommandGroupInterface
		Options          []*discordgo.ApplicationCommandOption
	}

	BaseSubCommandGroup struct {
		Name        string
		Description string
		SubCommands []SubCommandInterface
	}

	BaseSubCommand struct {
		Name        string
		Description string
		Options     []*discordgo.ApplicationCommandOption
	}
)

func (c *BaseCommand) GetName() string {
	return c.Name
}

func (c *BaseCommand) GetSubCommands() []SubCommandInterface {
	return c.SubCommands
}

func (c *BaseCommand) GetSubCommand(name string) (SubCommandInterface, bool) {
	for _, sub := range c.SubCommands {
		if sub.GetName() == name {
			return sub, true
		}
	}

	return nil, false
}

func (c *BaseCommand) GetSubCommandGroups() []SubCommandGroupInterface {
	return c.SubCommandGroups
}

func (c *BaseCommand) GetSubCommandGroup(name string) (SubCommandGroupInterface, bool) {
	for _, sub := range c.SubCommandGroups {
		if sub.GetName() == name {
			return sub, true
		}
	}

	return nil, false
}

func (c *BaseCommand) GetApplicationCommand() *discordgo.ApplicationCommand {
	ac := &discordgo.ApplicationCommand{
		Name:        c.Name,
		Description: c.Description,
		Type:        discordgo.ChatApplicationCommand,
		Options:     c.Options,
	}

	options := make([]*discordgo.ApplicationCommandOption, 0)

	for _, subcommand := range c.SubCommands {
		fmt.Println("Adding subcommand", subcommand.GetName())
		options = append(options, subcommand.GetApplicationCommandOption())
	}

	for _, subcommandGroup := range c.SubCommandGroups {
		fmt.Println("Adding subcommand group", subcommandGroup.GetName())
		options = append(options, subcommandGroup.GetApplicationCommandOption())
	}

	if len(options) > 0 {
		ac.Options = options
	}

	ac.Options = append(ac.Options, options...)

	return ac
}

func (c *BaseSubCommandGroup) GetName() string {
	return c.Name
}

func (c *BaseSubCommandGroup) GetSubCommands() []SubCommandInterface {
	return c.SubCommands
}

func (c *BaseSubCommandGroup) GetSubCommand(name string) (SubCommandInterface, bool) {
	for _, sub := range c.SubCommands {
		if sub.GetName() == name {
			return sub, true
		}
	}

	return nil, false
}

func (c *BaseSubCommandGroup) GetApplicationCommandOption() *discordgo.ApplicationCommandOption {
	options := make([]*discordgo.ApplicationCommandOption, 0)

	for _, subcommand := range c.SubCommands {
		fmt.Println("Adding subcommand", subcommand.GetName(), "to group", c.Name)
		options = append(options, subcommand.GetApplicationCommandOption())
	}

	return &discordgo.ApplicationCommandOption{
		Name:        c.Name,
		Description: c.Description,
		Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
		Options:     options,
	}
}

func (c *BaseSubCommand) GetName() string {
	return c.Name
}

func (c *BaseSubCommand) GetApplicationCommandOption() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name:        c.Name,
		Description: c.Description,
		Type:        discordgo.ApplicationCommandOptionSubCommand,
		Options:     c.Options,
	}
}
