package router

import (
	"fmt"

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
			return NewErrCommandAlreadyRegistered(command.Name)
		}

		fmt.Println("Registering command", command.Name)

		r.commands[command.Name] = command
	}

	return nil
}

func (r *CommandRouter) Get(name string) (*Command, error) {
	cmd, ok := r.commands[name]
	if !ok {
		return nil, NewErrCommandNotFound(name)
	}

	return cmd, nil

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
	if c.Commands != nil && len(c.Commands.List()) == 0 {
		options := make([]*discordgo.ApplicationCommandOption, len(c.Options))
		for i, option := range c.Options {
			options[i] = (*discordgo.ApplicationCommandOption)(option)
		}

		return options
	}

	options := make([]*discordgo.ApplicationCommandOption, len(c.Commands.List()))

	for i, command := range c.Commands.List() {
		oType := discordgo.ApplicationCommandOptionSubCommand
		if command.Commands != nil && len(command.Commands.List()) > 0 {
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
