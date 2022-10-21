package router

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type (
	MetaRouter interface {
		Register(commands ...*Command) error
		Unregister(commands ...*Command)
		Handler(s *discordgo.Session, i *discordgo.InteractionCreate)
		Sync(s *discordgo.Session, guildid string) error
	}

	Router struct {
		commands *CommandRouter
	}
)

func NewRouter() (*Router, error) {
	r := &Router{
		commands: NewCommandRouter(nil),
	}

	return r, nil
}

func (r *Router) RegisterCommands(commands ...*Command) error {
	return r.commands.Register(commands...)
}

func (r *Router) UnregisterCommands(commands ...*Command) {
	r.commands.Unregister(commands...)
}

// Handler is registered with discordgo to handle all interaction events
func (r *Router) Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionPing:
		go r.handlePing(s, i)
	case discordgo.InteractionApplicationCommand:
		go r.handleApplicationCommand(s, i)
	case discordgo.InteractionMessageComponent:
		go r.handleMessageComponent(s, i)
	case discordgo.InteractionApplicationCommandAutocomplete:
		go r.handleCommandAutocomplete(s, i)
	case discordgo.InteractionModalSubmit:
		go r.handleModalSubmit(s, i)
	}
}

// Sync syncs all commands and subcommands with discord
func (r *Router) Sync(s *discordgo.Session) error {
	cmds := make([]*discordgo.ApplicationCommand, 0, len(r.commands.List()))
	for i, c := range r.commands.List() {
		cmds[i] = c.ApplicationCommand()
	}

	appcmds, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", cmds)
	if err != nil {
		return err
	}

	for _, c := range appcmds {
		fmt.Println("Synced command", c.Name)
	}

	return nil

}

func (r *Router) handlePing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponsePong,
	})
}

func (r *Router) handleApplicationCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	cmd := r.commands.Get(data.Name)
	if cmd == nil {
		fmt.Println("Command not found", data.Name)
		return
	}

	handler := cmd.Handler
	opts := data.Options

	if len(data.Options) > 0 {
		switch data.Options[0].Type {
		case discordgo.ApplicationCommandOptionSubCommand:
			subcmd := cmd.Commands.Get(data.Options[0].Name)
			if subcmd == nil {
				fmt.Println("Subcommand not found", data.Options[0].Name)
				return
			}

			handler = subcmd.Handler
			opts = data.Options[0].Options
		case discordgo.ApplicationCommandOptionSubCommandGroup:
			group := cmd.Commands.Get(data.Options[0].Name)
			if group == nil {
				fmt.Println("Subcommand not found", data.Options[0].Name)
				return
			}

			if group.Commands == nil {
				fmt.Println("Subcommand group not found", data.Options[0].Name)
				return
			}

			subcmd := group.Commands.Get(data.Options[0].Options[0].Name)
			if subcmd == nil {
				fmt.Println("Subcommand not found", data.Options[0].Options[0].Name)
				return
			}

			handler = subcmd.Handler
			opts = data.Options[0].Options[0].Options
		}
	}

	ctx := &CommandContext{
		Context: &Context{
			Session:     s,
			Router:      r,
			Interaction: i.Interaction,
		},
		Options: convertOptionsToMap(opts),
	}

	go handler(ctx)
}

func (r *Router) handleMessageComponent(s *discordgo.Session, i *discordgo.InteractionCreate) {
}

func (r *Router) handleCommandAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// todo
}

func (r *Router) handleModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {

}
