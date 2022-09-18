package router

import (
	"github.com/bwmarrin/discordgo"
)

type (
	HandlerFunc func(s *discordgo.Session, i *discordgo.InteractionCreate)

	Router interface {
		// Discordgo interaction handler
		Handler() HandlerFunc
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
func (r *router) Handler() HandlerFunc {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		data := i.ApplicationCommandData()

		if h, ok := r.commands[data.Name]; ok {
			ctx := &Context{
				Session:     s,
				Interaction: i,
			}

			// Handle subcommands
			if len(data.Options) != 0 {
				if sub, ok := h.GetSubCommand(data.Options[0].Name); ok {
					if len(data.Options[0].Options) != 0 {
						if sub, ok = sub.GetSubCommand(data.Options[0].Options[0].Name); ok {
							h = sub
						} else {
							return
						}
					} else {
						h = sub
					}
				} else {
					return
				}
			}

			h.Handler(ctx)

		}
	}
}

// Sync syncs all commands and subcommands with discord
func (r *router) Sync(s *discordgo.Session) error {
	commands := make([]*discordgo.ApplicationCommand, 0)
	for _, c := range r.commands {
		commands = append(commands, c.GetApplicationCommand())
	}

	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", commands)
	return err
}

func (r *router) RegisterCommands(commands ...CommandInterface) {
	for _, c := range commands {
		r.commands[c.GetName()] = c
	}
}
