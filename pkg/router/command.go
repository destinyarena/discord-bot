package router

import "github.com/bwmarrin/discordgo"

type (
	HandlerFunc      func(ctx *Context)
	CommandOption    discordgo.ApplicationCommandOption
	CommandInterface interface {
		GetName() string
		GetDescription() string
		GetHandler() HandlerFunc
		GetSubCommands() []*SubCommand
		GetSubCommandGroups() []*SubCommandGroup
		GetOptions() []*CommandOption
		GetDMPermission() bool
		GetDefaultPermmissions() uint64
		ToApplicationCommand() *discordgo.ApplicationCommand
	}

	SubCommandGroupInterface interface {
		GetName() string
		GetDescription() string
		GetSubCommands() []*SubCommand
	}

	SubCommandInterface interface {
		GetName() string
		GetDescription() string
		GetHandler() HandlerFunc
		GetOptions() []*CommandOption
	}

	Command struct {
		Name                string
		Description         string
		SubCommandGroups    map[string]SubCommandGroupInterface
		SubCommands         map[string]SubCommandInterface
		Options             []*CommandOption
		Handler             HandlerFunc
		DefaultPermmissions *int64
		DMPermmission       *bool
	}

	SubCommandGroup struct {
		Name        string
		Description string
		SubCommands map[string]SubCommandInterface
	}

	SubCommand struct {
		Name        string
		Description string
		Options     []*CommandOption
		Handler     HandlerFunc
	}
)

func NewCommand(name string, description string, handler HandlerFunc) *Command {
	return &Command{
		Name:             name,
		Description:      description,
		Handler:          handler,
		SubCommands:      make(map[string]SubCommandInterface),
		SubCommandGroups: make(map[string]SubCommandGroupInterface),
		Options:          make([]*CommandOption, 0),
	}
}

func (c *Command) ToApplicationCommand() *discordgo.ApplicationCommand {
	options := make([]*discordgo.ApplicationCommandOption, 0)

	if len(c.SubCommands) > 0 || len(c.SubCommandGroups) > 0 {
		for _, g := range c.SubCommandGroups {
			goptions := make([]*discordgo.ApplicationCommandOption, 0)

			for _, s := range g.SubCommands {
				goptions = append(goptions, &discordgo.ApplicationCommandOption{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        s.Name,
					Description: s.Description,
					Options:     convertOptions(s.Options),
				})
			}

			options = append(options, &discordgo.ApplicationCommandOption{
				Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
				Name:        g.Name,
				Description: g.Description,
				Options:     goptions,
			})
		}

		for _, c := range c.SubCommands {
			options = append(options, &discordgo.ApplicationCommandOption{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        c.Name,
				Description: c.Description,
				Options:     convertOptions(c.Options),
			})
		}

	} else {
		options = append(options, convertOptions(c.Options)...)
	}

	return &discordgo.ApplicationCommand{
		Name:                     c.Name,
		Description:              c.Description,
		DefaultMemberPermissions: c.DefaultPermmissions,
		DMPermission:             c.DMPermmission,
		Options:                  options,
	}
}

func (c *Command) AddSubCommandGroup(name, description string) *Group {
	if c.SubCommandGroups == nil {
		c.SubCommandGroups = make(map[string]*Group)
	}

	c.SubCommandGroups[name] = &Group{
		Name:        name,
		Description: description,
		SubCommands: make(map[string]*Command),
	}

	return c.SubCommandGroups[name]
}

func (c *Command) GetSubCommandGroup(name string) *Group {
	if c.SubCommandGroups == nil {
		return nil
	}

	return c.SubCommandGroups[name]
}

func (c *Command) AddSubCommand(name, description string, options []*CommandOption, handler HandlerFunc) {
	if c.SubCommands == nil {
		c.SubCommands = make(map[string]*Command)
	}

	c.SubCommands[name] = &Command{
		Name:        name,
		Description: description,
		Options:     options,
		Handler:     handler,
	}
}

func (c *Command) GetSubCommand(name string) *Command {
	if c.SubCommands == nil {
		return nil
	}

	return c.SubCommands[name]
}

func (g *Group) AddSubCommand(name, description string, options []*CommandOption, handler HandlerFunc) {
	if g.SubCommands == nil {
		g.SubCommands = make(map[string]*Command)
	}

	g.SubCommands[name] = &Command{
		Name:        name,
		Description: description,
		Options:     options,
		Handler:     handler,
	}
}

func (g *Group) GetSubCommand(name string) *Command {
	if g.SubCommands == nil {
		return nil
	}

	return g.SubCommands[name]
}
