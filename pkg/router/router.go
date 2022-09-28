package router

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type (
	RouterInterface interface {
		// Discordgo interaction handler
		Handler(s *discordgo.Session, i *discordgo.InteractionCreate)

		// PreRoute Handlers
		AddPreRoute(h ...PreRouteInterface) (*PreRouteContext, error)

		// Syncs slash commands with discord
		Sync(s *discordgo.Session, guildid string) error

		// Registers Global Commands
		AddCommands(commands ...*Command) error

		// Updates registered commands
		UpdateCommands(commands ...*Command) error

		// Adds Guild Commands
		AddGuildCommands(guildid string, commands ...*Command) error
		UpdateGuildCommands(guildid string, commands ...*Command) error

		// Registers Button
		AddComponents(components ...*Component) error

		// Register Modal
		AddModals(modals ...*Modal) error

		// Registers Modules
		RegisterModules(m ...Module) error
	}
	Router struct {
		commands   map[string]*Command
		components map[string]*Component
		guilds     map[string]map[string]*Command
	}
)

func NewRouter() (Router, error) {
	r := &Router{
		commands:   make(map[string]*Command),
		components: make(map[string]*Component),
	}

	return r, nil
}

// Handler is registered with discordgo to handle all interaction events
func (r *Router) Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
func (r *Router) Sync(s *discordgo.Session, guildid string) error {
	var appcmds []*discordgo.ApplicationCommand

	if guildid != "" {
		cmds, ok := r.guilds[guildid]
		if !ok {
			return fmt.Errorf("no commands registered for guild: %s", guildid)
		}
		for _, c := range cmds {
			appcmds = append(appcmds, c.ToApplicationCommand())
		}
	} else {
		for _, c := range r.commands {
			appcmds = append(appcmds, c.ToApplicationCommand())
		}
	}

	uappcmds, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, guildid, appcmds)
	if err != nil {
		return err
	}

	for _, c := range uappcmds {
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

func (r *Router) handleMessageComponent(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.MessageComponentData()
	fmt.Printf("Received component: %s\n", data.CustomID)

	slice := strings.Split(data.CustomID, "-")
	if component, ok := r.components[slice[0]]; ok && (len(component.Args) == (len(slice) - 1)) {
		fmt.Printf("Found component: %s\n", slice[0])

		args := make(map[string]*ComponentArgument, 0)

		for idx, arg := range component.Args {
			var value interface{}

			fmt.Println("Args for component:", component.Name, slice[1:])

			switch arg.Type {
			case ComponentArgumentTypeString:
				value = string(slice[idx+1])
			case ComponentArgumentTypeUser:
				value, _ = s.User(slice[idx+1])
			case ComponentArgumentTypeChannel:
				value, _ = s.Channel(slice[idx+1])
			case ComponentArgumentTypeRole:
				value, _ = s.State.Role(i.GuildID, slice[idx+1])
			}

			args[arg.Name] = &ComponentArgument{
				Name:  arg.Name,
				Value: value,
				Type:  arg.Type,
			}
		}

		ctx := &ComponentContext{
			Context: Context{
				Session:     s,
				Interaction: i.Interaction,
				Message:     i.Message,
			},
			CustomID: data.CustomID,
			Args:     args,
		}

		go component.Handler(ctx)

	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "Invalid component",
			},
		})

	}
}

func (r *Router) handleCommandAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// todo
}
