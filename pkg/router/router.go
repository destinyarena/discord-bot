package router

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type (
	HandlerFunc func(ctx *Context)
	Module      interface {
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
		// Registers component handler to the router
		RegisterComponents(components ...*Component)
		// Registers modules which may include commands and components
		RegisterModules(m ...Module)
	}
	router struct {
		commands   map[string]*Command
		components map[string]*Component
		modals     map[string]*Modal
	}
)

func New() (Router, error) {
	r := &router{
		commands:   make(map[string]*Command),
		components: make(map[string]*Component),
		modals:     make(map[string]*Modal),
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
	case discordgo.InteractionPing:
		r.handlePing(s, i)
	case discordgo.InteractionApplicationCommand:
		r.handleApplicationCommand(s, i)
	case discordgo.InteractionMessageComponent:
		r.handleMessageComponent(s, i)
	case discordgo.InteractionApplicationCommandAutocomplete:
		r.handleCommandAutocomplete(s, i)
	case discordgo.InteractionModalSubmit:
		//r.handleModalSubmit(s, i)
		fmt.Println("TODO")
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

func (r *router) RegisterComponents(components ...*Component) {
	for _, c := range components {
		r.components[c.Name] = c
	}
}

func convertOptionsToMap(options []*discordgo.ApplicationCommandInteractionDataOption) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	m := make(map[string]*discordgo.ApplicationCommandInteractionDataOption)
	for _, o := range options {
		m[o.Name] = o
	}

	return m
}

func (r *router) handlePing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponsePong,
	})
}

func (r *router) handleApplicationCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	c, ok := r.commands[data.Name]
	if !ok {
		fmt.Printf("Command not found: %s\n", data.Name)
		return
	}

	ctx := &Context{
		Session:     s,
		Interaction: i.Interaction,
		Options:     make(map[string]*discordgo.ApplicationCommandInteractionDataOption),
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

func (r *router) handleMessageComponent(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.MessageComponentData()
	fmt.Printf("Received component: %s\n", data.CustomID)

	if c, ok := r.components[data.CustomID]; ok {
		ctx := &ComponentContext{
			Session:     s,
			Interaction: i.Interaction,
			Message:     i.Message,
		}

		go c.Handler(ctx)
	}
}

func (r *router) handleCommandAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// todo
}
