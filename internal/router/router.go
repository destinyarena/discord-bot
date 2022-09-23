package router

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type (
	Router interface {
		// Discordgo interaction handler
		Handler(s *discordgo.Session, i *discordgo.InteractionCreate)
		// Syncs slash commands with discord
		Sync(s *discordgo.Session) error
		// Registers commands to the router
		RegisterCommands(commands ...CommandInterface)
	}

	router struct {
		commands map[string]CommandInterface
	}
)

func New() (Router, error) {
	r := &router{
		commands: make(map[string]CommandInterface),
	}

	return r, nil
}

// Handler is registered with discordgo to handle all interaction events
func (r *router) Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()
	fmt.Printf("Got command %s\n", data.Name)

	if h, ok := r.commands[data.Name]; ok {

		ctx := &Context{
			Session:           s,
			Interaction:       i.Interaction,
			InteractionCreate: i,
			Options:           make(map[string]*discordgo.ApplicationCommandInteractionDataOption),
		}

		// Handle subcommands
		if len(data.Options) != 0 {
			switch data.Options[0].Type {
			case discordgo.ApplicationCommandOptionSubCommand:
				if sub, ok := h.GetSubCommand(data.Options[0].Name); ok {
					for _, o := range data.Options[0].Options {
						ctx.Options[o.Name] = o
					}
					go sub.Handler(ctx)
					return
				}
			case discordgo.ApplicationCommandOptionSubCommandGroup:
				if group, ok := h.GetSubCommandGroup(data.Options[0].Name); ok {
					if sub, ok := group.GetSubCommand(data.Options[0].Options[0].Name); ok {
						for _, o := range data.Options[0].Options[0].Options {
							ctx.Options[o.Name] = o
						}
						go sub.Handler(ctx)
						return
					}
				}
			}
		}

		for _, o := range data.Options {
			ctx.Options[o.Name] = o
		}

		go h.Handler(ctx)
	}
}

// Sync syncs all commands and subcommands with discord
func (r *router) Sync(s *discordgo.Session) error {
	commands := make([]*discordgo.ApplicationCommand, 0)
	for _, c := range r.commands {
		commands = append(commands, c.GetApplicationCommand())
	}

	cmds, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", commands)
	if err != nil {
		return err
	}

	for _, c := range cmds {
		fmt.Println("Registered command", c.Name)
	}

	return nil

}

func (r *router) RegisterCommands(commands ...CommandInterface) {
	for _, c := range commands {
		r.commands[c.GetName()] = c
	}
}
