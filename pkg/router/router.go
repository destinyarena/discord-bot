package router

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type (
	HandlerFunc func(ctx *Context)

	Component struct{}
	Module    interface {
		Commands() []*Command
		Components() []*Component
	}
	Router interface {
		// Discordgo interaction handler
		Handler(s *discordgo.Session, i *discordgo.InteractionCreate)
		// Syncs slash commands with discord
		Sync(s *discordgo.Session, guildid string) error
		// Registers commands to the router
		RegisterCommands(commands ...*Command)
		// Registers modules which may include commands and components
		RegisterModules(m ...Module)
	}
	router struct {
		commands map[string]*Command
	}
)

func New() (Router, error) {
	r := &router{
		commands: make(map[string]*Command),
	}

	return r, nil
}

func (r *router) RegisterModules(m ...Module) {
	for _, mod := range m {
		r.RegisterCommands(mod.Commands()...)
	}
}

// Handler is registered with discordgo to handle all interaction events
func (r *router) Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		r.handleCommand(s, i)
	case discordgo.InteractionMessageComponent:
		r.handleComponent(s, i)
	default:
		fmt.Printf("Unknown interaction type: %d\n", i.Type)
	}
}

// Sync syncs all commands and subcommands with discord
func (r *router) Sync(s *discordgo.Session, guildid string) error {
	commands := make([]*discordgo.ApplicationCommand, 0)
	for _, c := range r.commands {
		commands = append(commands, c.ApplicationCommand())
	}

	cmds, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, guildid, commands)
	if err != nil {
		return err
	}

	for _, c := range cmds {
		fmt.Println("Registered command", c.Name)
	}

	return nil

}

func (r *router) RegisterCommands(commands ...*Command) {

	for _, c := range commands {
		r.commands[c.Name] = c
	}
}

func convertOptionsToMap(options []*discordgo.ApplicationCommandInteractionDataOption) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	m := make(map[string]*discordgo.ApplicationCommandInteractionDataOption)
	for _, o := range options {
		m[o.Name] = o
	}

	return m
}

func (r *router) handleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	c, ok := r.commands[data.Name]
	if !ok {
		fmt.Printf("Command not found: %s\n", data.Name)
		return
	}

	ctx := &Context{
		Session:           s,
		Interaction:       i.Interaction,
		InteractionCreate: i,
		Options:           make(map[string]*discordgo.ApplicationCommandInteractionDataOption),
	}

	var handler HandlerFunc

	if len(data.Options) > 0 {
		switch data.Options[0].Type {
		case discordgo.ApplicationCommandOptionSubCommand:
			if sub := c.GetSubCommand(data.Options[0].Name); sub != nil {
				handler = sub.Handler
				ctx.Options = convertOptionsToMap(data.Options[0].Options)
			}
		case discordgo.ApplicationCommandOptionSubCommandGroup:
			if g := c.GetSubCommandGroup(data.Options[0].Name); g != nil {
				if sub := g.GetSubCommand(data.Options[0].Options[0].Name); sub != nil {
					handler = sub.Handler
					ctx.Options = convertOptionsToMap(data.Options[0].Options[0].Options)
				}
			}
		default:
			ctx.Options = convertOptionsToMap(data.Options)
		}
	}

	go handler(ctx)
}

func (r *router) handleComponent(s *discordgo.Session, i *discordgo.InteractionCreate) {

}
