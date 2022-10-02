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

	CommandInterface interface {
		GetName() string
		GetDescription() string
		GetHandler() CommandHandlerFunc
		AddSubCommands(subcommands ...SubCommandInterface) error
		GetSubCommand(name string) SubCommandInterface
		GetSubCommands() []SubCommandInterface
		AddSubCommandGroups(subcommandgroups ...SubCommandGroupInterface) error
		GetSubCommandGroup(name string) SubCommandGroupInterface
		GetSubCommandGroups() []SubCommandGroupInterface
		AddComponents(components ...*Component) error
		GetComponents() []*Component
		AddModals(modals ...*Modal) error
		GetModals() []*Modal
		GetOptions() []*CommandOption
		GetDMPermission() bool
		GetDefaultPermmissions() int64
		ToApplicationCommand() *discordgo.ApplicationCommand
	}

	SubCommandGroupInterface interface {
		GetName() string
		GetDescription() string
		AddSubCommands(subcommands ...SubCommandInterface) error
		GetSubCommand(name string) SubCommandInterface
		GetSubCommands() []SubCommandInterface
	}

	SubCommandInterface interface {
		GetName() string
		GetDescription() string
		GetHandler() CommandHandlerFunc
		GetOptions() []*CommandOption
	}

	Command struct {
		Name                string
		Description         string
		SubCommandGroups    map[string]SubCommandGroupInterface
		SubCommands         map[string]SubCommandInterface
		Modals              []*Modal
		Components          []*Component
		Options             []*CommandOption
		Handler             CommandHandlerFunc
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
		Handler     CommandHandlerFunc
	}
)

func NewCommand(name string, description string, handler CommandHandlerFunc) *Command {
	return &Command{
		Name:             name,
		Description:      description,
		Handler:          handler,
		SubCommands:      make(map[string]SubCommandInterface),
		SubCommandGroups: make(map[string]SubCommandGroupInterface),
		Components:       make([]*Component, 0),
		Modals:           make([]*Modal, 0),
		Options:          make([]*CommandOption, 0),
	}
}

func (c *Command) GetName() string {
	return c.Name
}

func (c *Command) GetDescription() string {
	return c.Description
}

func (c *Command) GetHandler() CommandHandlerFunc {
	return c.Handler
}

func (c *Command) AddSubCommands(subcommands ...SubCommandInterface) error {
	if c.SubCommands == nil {
		c.SubCommands = make(map[string]SubCommandInterface)
	}

	for _, s := range subcommands {
		if _, ok := c.SubCommands[s.GetName()]; ok {
			return fmt.Errorf("subcommand already exists: %s", s.GetName())
		}
		c.SubCommands[s.GetName()] = s
	}
	return nil
}

func (c *Command) GetSubCommand(name string) SubCommandInterface {
	return c.SubCommands[name]
}

func (c *Command) GetSubCommands() []SubCommandInterface {
	var subCommands []SubCommandInterface
	for _, subCommand := range c.SubCommands {
		subCommands = append(subCommands, subCommand)
	}
	return subCommands
}

func (c *Command) AddSubCommandGroups(subcommandgroups ...SubCommandGroupInterface) error {
	if c.SubCommandGroups == nil {
		c.SubCommandGroups = make(map[string]SubCommandGroupInterface)
	}

	for _, s := range subcommandgroups {
		if _, ok := c.SubCommandGroups[s.GetName()]; ok {
			return fmt.Errorf("subcommandgroup already exists: %s", s.GetName())
		}
		c.SubCommandGroups[s.GetName()] = s
	}
	return nil
}

func (c *Command) GetSubCommandGroup(name string) SubCommandGroupInterface {
	return c.SubCommandGroups[name]
}

func (c *Command) GetSubCommandGroups() []SubCommandGroupInterface {
	var subCommandGroups []SubCommandGroupInterface
	for _, subCommandGroup := range c.SubCommandGroups {
		subCommandGroups = append(subCommandGroups, subCommandGroup)
	}
	return subCommandGroups
}

func (c *Command) AddComponents(components ...*Component) error {
	c.Components = append(c.Components, components...)
	return nil
}

func (c *Command) GetComponents() []*Component {
	return c.Components
}

func (c *Command) AddModals(modals ...*Modal) error {
	c.Modals = append(c.Modals, modals...)
	return nil
}

func (c *Command) GetModals() []*Modal {
	return c.Modals
}

func (c *Command) GetOptions() []*CommandOption {
	return c.Options
}

func (c *Command) GetDMPermission() bool {
	if c.DMPermmission != nil {
		return *c.DMPermmission
	}
	return false
}

func (c *Command) GetDefaultPermmissions() int64 {
	if c.DefaultPermmissions != nil {
		return int64(*c.DefaultPermmissions)
	}
	return 0
}

func (c *Command) ToApplicationCommand() *discordgo.ApplicationCommand {
	options := make([]*discordgo.ApplicationCommandOption, 0)

	if len(c.SubCommands) > 0 || len(c.SubCommandGroups) > 0 {
		for _, g := range c.SubCommandGroups {
			goptions := make([]*discordgo.ApplicationCommandOption, 0)

			for _, sc := range g.GetSubCommands() {
				goptions = append(goptions, &discordgo.ApplicationCommandOption{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        sc.GetName(),
					Description: sc.GetDescription(),
					Options:     convertOptions(sc.GetOptions()),
				})
			}

			options = append(options, &discordgo.ApplicationCommandOption{
				Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
				Name:        g.GetName(),
				Description: g.GetDescription(),
				Options:     goptions,
			})
		}

		for _, sc := range c.SubCommands {
			options = append(options, &discordgo.ApplicationCommandOption{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        sc.GetName(),
				Description: sc.GetDescription(),
				Options:     convertOptions(sc.GetOptions()),
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

func (scg *SubCommandGroup) GetName() string {
	return scg.Name
}

func (scg *SubCommandGroup) GetDescription() string {
	return scg.Description
}

func (scg *SubCommandGroup) AddSubCommands(subcommands ...SubCommandInterface) error {
	if scg.SubCommands == nil {
		scg.SubCommands = make(map[string]SubCommandInterface)
	}

	for _, s := range subcommands {
		if _, ok := scg.SubCommands[s.GetName()]; ok {
			return fmt.Errorf("subcommand already exists: %s", s.GetName())
		}
		scg.SubCommands[s.GetName()] = s
	}
	return nil
}

func (scg *SubCommandGroup) GetSubCommand(name string) SubCommandInterface {
	return scg.SubCommands[name]
}

func (scg *SubCommandGroup) GetSubCommands() []SubCommandInterface {
	var subCommands []SubCommandInterface
	for _, subCommand := range scg.SubCommands {
		subCommands = append(subCommands, subCommand)
	}
	return subCommands
}

func (sc *SubCommand) GetName() string {
	return sc.Name
}

func (sc *SubCommand) GetDescription() string {
	return sc.Description
}

func (sc *SubCommand) GetHandler() CommandHandlerFunc {
	return sc.Handler
}

func (sc *SubCommand) GetOptions() []*CommandOption {
	return sc.Options
}
