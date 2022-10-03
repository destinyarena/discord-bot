package router

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type (
	Router struct {
		commands      map[string]CommandInterface
		guildCommands map[string]map[string]CommandInterface
		components    map[string]*Component
		modals        map[string]*Modal
	}
)

func NewRouter() (*Router, error) {
	r := &Router{
		commands:      make(map[string]CommandInterface),
		guildCommands: make(map[string]map[string]CommandInterface),
		components:    make(map[string]*Component),
		modals:        make(map[string]*Modal),
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
		r.handleModalSubmit(s, i)
	}
}

// Sync syncs all commands and subcommands with discord
func (r *Router) Sync(s *discordgo.Session, guildid string) error {
	var appcmds []*discordgo.ApplicationCommand

	if guildid != "" {
		cmds, ok := r.guildCommands[guildid]
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

func (r *Router) AddCommands(commands ...CommandInterface) error {
	for _, c := range commands {
		if _, ok := r.commands[c.GetName()]; ok {
			return &CommandExistsError{c.GetName()}
		}
		r.commands[c.GetName()] = c
		fmt.Printf("Adding command %s\n", c.GetName())
		r.AddComponents(c.GetComponents()...)
		r.AddModals(c.GetModals()...)
	}
	return nil
}

func (r *Router) UpdateCommands(commands ...CommandInterface) error {
	for _, c := range commands {
		if _, ok := r.commands[c.GetName()]; !ok {
			return &CommandNotFoundError{c.GetName()}
		}
		r.commands[c.GetName()] = c
	}
	return nil
}

func (r *Router) AddGuildCommands(guildid string, commands ...CommandInterface) error {
	if _, ok := r.guildCommands[guildid]; !ok {
		r.guildCommands[guildid] = make(map[string]CommandInterface)
	}

	for _, c := range commands {
		if _, ok := r.guildCommands[guildid][c.GetName()]; ok {
			return &CommandExistsError{c.GetName()}
		}
		r.guildCommands[guildid][c.GetName()] = c
		fmt.Printf("Adding command %s to guild %s\n", c.GetName(), guildid)
	}

	return nil
}

func (r *Router) UpdateGuildCommands(guildid string, commands ...CommandInterface) error {
	if _, ok := r.guildCommands[guildid]; !ok {
		return fmt.Errorf("no commands registered for guild: %s", guildid)
	}

	for _, c := range commands {
		if _, ok := r.guildCommands[guildid][c.GetName()]; !ok {
			return &CommandNotFoundError{c.GetName()}
		}

		r.guildCommands[guildid][c.GetName()] = c
	}

	return nil
}

func (r *Router) AddComponents(components ...*Component) error {
	for _, c := range components {
		if _, ok := r.components[c.ID]; ok {
			return &ComponentExistsError{c.ID}
		}
		r.components[c.ID] = c
		fmt.Printf("Adding component %s\n", c.ID)
	}
	return nil
}

func (r *Router) GetComponent(id string) *Component {
	return r.components[id]
}

func (r *Router) BuildComponent(id string, args ...interface{}) (discordgo.MessageComponent, error) {
	c, ok := r.components[id]
	if !ok {
		return nil, &ComponentNotFoundError{id}
	}

	return c.Build(args...)
}

func (r *Router) AddModals(modals ...*Modal) error {
	for _, m := range modals {
		if _, ok := r.modals[m.ID]; ok {
			return &ModalExistsError{m.ID}
		}
		r.modals[m.ID] = m
	}
	return nil
}

func (r *Router) GetModal(id string) *Modal {
	return r.modals[id]
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
		err := &CommandNotFoundError{data.Name}
		fmt.Println(err)
		return
	}

	ctx := &CommandContext{
		Context: &Context{
			Session:     s,
			Router:      r,
			Interaction: i.Interaction,
		},
		Options: make(map[string]*discordgo.ApplicationCommandInteractionDataOption),
	}

	var handler CommandHandlerFunc

	if len(data.Options) > 0 {
		switch data.Options[0].Type {
		case discordgo.ApplicationCommandOptionSubCommand:
			if sub := c.GetSubCommand(data.Options[0].Name); sub != nil {
				handler = sub.GetHandler()
				ctx.Options = convertOptionsToMap(data.Options[0].Options)
			}
		case discordgo.ApplicationCommandOptionSubCommandGroup:
			if g := c.GetSubCommandGroup(data.Options[0].Name); g != nil {
				if sub := g.GetSubCommand(data.Options[0].Options[0].Name); sub != nil {
					handler = sub.GetHandler()
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
	if component, ok := r.components[slice[0]]; ok && (len(component.Args) == (len(slice) - 1)) && component.Type == data.ComponentType {
		fmt.Printf("Found component: %s\n", slice[0])

		args := make(map[string]*ComponentArgument, 0)

		for idx, arg := range component.Args {
			var value interface{}

			fmt.Println("Args for component:", component.ID, slice[1:])

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
				Router:      r,
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

func (r *Router) handleModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {

}
