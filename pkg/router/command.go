package router

import (
	"github.com/bwmarrin/discordgo"
)

type (
	CommandHandlerFunc func(ctx *CommandContext)
	CommandOption      discordgo.ApplicationCommandOption

	CommandContext struct {
		*Context
		Options map[string]*discordgo.ApplicationCommandInteractionDataOption
	}

	CommandRouter struct {
		commands map[string]*Command
	}

	Command struct {
		Name                string
		Description         string
		Commands            *CommandRouter
		Options             []*CommandOption
		Handler             CommandHandlerFunc
		DefaultPermmissions *int64
		DMPermmission       *bool
	}
)

func NewCommandRouter(commands []*Command) *CommandRouter {
	router := &CommandRouter{
		commands: make(map[string]*Command),
	}

	for _, command := range commands {
		router.Register(command)
	}

	return router
}

func (r *CommandRouter) Register(commands ...*Command) error {
	for _, command := range commands {
		if _, ok := r.commands[command.Name]; ok {
			return &CommandAlreadyRegisteredError{command.Name}
		}

		r.commands[command.Name] = command
	}

	return nil
}

func (r *CommandRouter) Get(name string) *Command {
	return r.commands[name]
}

func (r *CommandRouter) Unregister(commands ...*Command) error {
	for _, command := range commands {
		delete(r.commands, command.Name)
	}

	return nil
}

func (r *CommandRouter) List() []*Command {
	commands := make([]*Command, len(r.commands))

	i := 0
	for _, command := range r.commands {
		commands[i] = command
		i++
	}

	return commands
}

func (c *Command) applicationCommandOptions() []*discordgo.ApplicationCommandOption {
	if c.Commands == nil {
		options := make([]*discordgo.ApplicationCommandOption, len(c.Options))
		for i, option := range c.Options {
			options[i] = (*discordgo.ApplicationCommandOption)(option)
		}

		return options
	}

	options := make([]*discordgo.ApplicationCommandOption, len(c.Commands.List()))

	for i, command := range c.Commands.List() {
		oType := discordgo.ApplicationCommandOptionSubCommand
		if command.Commands != nil {
			oType = discordgo.ApplicationCommandOptionSubCommandGroup
		}

		options[i] = &discordgo.ApplicationCommandOption{
			Type:        oType,
			Name:        command.Name,
			Description: command.Description,
			Options:     command.applicationCommandOptions(),
		}
	}

	return options
}

func (c *Command) ApplicationCommand() *discordgo.ApplicationCommand {
	ac := &discordgo.ApplicationCommand{
		Name:                     c.Name,
		Type:                     discordgo.ChatApplicationCommand,
		Description:              c.Description,
		DMPermission:             c.DMPermmission,
		DefaultMemberPermissions: c.DefaultPermmissions,
	}

	ac.Options = c.applicationCommandOptions()

	return ac
}

// Command Builder

func NewCommand(name, description string) *Command {
	return &Command{
		Name:        name,
		Description: description,
		Commands:    NewCommandRouter(nil),
	}
}

func (c *Command) WithOptions(options ...*CommandOption) *Command {
	c.Options = options

	return c
}

func (c *Command) WithHandler(handler CommandHandlerFunc) *Command {
	c.Handler = handler

	return c
}

func (c *Command) WithDefaultPermissions(permissions int64) *Command {
	c.DefaultPermmissions = &permissions

	return c
}

func (c *Command) WithDMPermissions(allowed bool) *Command {
	c.DMPermmission = &allowed

	return c
}

func (c *Command) WithSubCommands(commands ...*Command) *Command {
	c.Commands.Register(commands...)

	return c
}
