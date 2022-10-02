package router

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type (
	ComponentHandlerFunc func(ctx *ComponentContext)
	ComponentContext     struct {
		Context
		CustomID string
		Args     map[string]*ComponentArgument
	}

	ComponentArgumentType int
	ComponentArgument     struct {
		Name  string
		Value interface{}
		Type  ComponentArgumentType
	}

	Component struct {
		ID       string
		Type     discordgo.ComponentType
		Args     []*ComponentArgument
		Disabled bool
		Handler  ComponentHandlerFunc

		Label string
		Style discordgo.ButtonStyle
		Emoji discordgo.ComponentEmoji

		Placeholder string
		MinValues   *int
		MaxValues   int
		Options     []discordgo.SelectMenuOption
	}
)

const (
	ComponentArgumentTypeString ComponentArgumentType = iota
	ComponentArgumentTypeUser
	ComponentArgumentTypeChannel
	ComponentArgumentTypeRole
)

func (c *Component) Build(args ...interface{}) (discordgo.MessageComponent, error) {
	if len(args) != len(c.Args) {
		return nil, fmt.Errorf("invalid number of arguments for component %s", c.ID)
	}

	customID := c.ID

	for idx, arg := range c.Args {
		switch arg.Type {
		case ComponentArgumentTypeString:
			value, ok := args[idx].(string)
			if !ok {
				return nil, fmt.Errorf("invalid argument type for component %s", c.ID)
			}
			customID += "-" + value
		case ComponentArgumentTypeUser:
			value, ok := args[idx].(*discordgo.User)
			if !ok {
				return nil, fmt.Errorf("invalid argument type for component %s", c.ID)
			}
			customID += "-" + value.ID
		case ComponentArgumentTypeChannel:
			value, ok := args[idx].(*discordgo.Channel)
			if !ok {
				return nil, fmt.Errorf("invalid argument type for component %s", c.ID)
			}
			customID += "-" + value.ID
		case ComponentArgumentTypeRole:
			value, ok := args[idx].(*discordgo.Role)
			if !ok {
				return nil, fmt.Errorf("invalid argument type for component %s", c.ID)
			}
			customID += "-" + value.ID
		}
	}

	switch c.Type {
	case discordgo.ButtonComponent:
		return &discordgo.Button{
			Label:    c.Label,
			Style:    c.Style,
			Emoji:    c.Emoji,
			CustomID: customID,
			Disabled: c.Disabled,
		}, nil
	case discordgo.SelectMenuComponent:
		return &discordgo.SelectMenu{
			CustomID:    customID,
			Placeholder: c.Placeholder,
			MinValues:   c.MinValues,
			MaxValues:   c.MaxValues,
			Options:     c.Options,
			Disabled:    c.Disabled,
		}, nil
	}

	return nil, fmt.Errorf("invalid component type for component %s", c.ID)
}
