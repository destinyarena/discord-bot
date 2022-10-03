package router

import (
	"errors"
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

	Component interface {
		GetType() discordgo.ComponentType
		Build(args ...interface{}) (discordgo.MessageComponent, error)
	}

	ButtonComponent struct {
		Component
		ID       string
		Label    string
		Style    discordgo.ButtonStyle
		Emoji    discordgo.ComponentEmoji
		Disabled bool
		Args     []*ComponentArgument
		Handler  ComponentHandlerFunc
	}

	SelectMenuComponent struct {
		Component
		ID          string
		Placeholder string
		MinValues   *int
		MaxValues   int
		Disabled    bool
		Options     []discordgo.SelectMenuOption
		Args        []*ComponentArgument
		Handler     ComponentHandlerFunc
	}
)

const (
	ComponentArgumentTypeString ComponentArgumentType = iota
	ComponentArgumentTypeUser
	ComponentArgumentTypeChannel
	ComponentArgumentTypeRole
)

func getIDFromArgs(baseID string, cargs []*ComponentArgument, args []interface{}) (string, error) {
	if len(args) != len(cargs) {
		return baseID, errors.New("Invalid number of component arguments")
	}

	customID := baseID

	for idx, arg := range cargs {
		switch arg.Type {
		case ComponentArgumentTypeString:
			value, ok := args[idx].(string)
			if !ok {
				return "", fmt.Errorf("invalid argument type for component %s", baseID)
			}
			customID += "-" + value
		case ComponentArgumentTypeUser:
			value, ok := args[idx].(*discordgo.User)
			if !ok {
				return "", fmt.Errorf("invalid argument type for component %s", baseID)
			}
			customID += "-" + value.ID
		case ComponentArgumentTypeChannel:
			value, ok := args[idx].(*discordgo.Channel)
			if !ok {
				return "", fmt.Errorf("invalid argument type for component %s", baseID)
			}
			customID += "-" + value.ID
		case ComponentArgumentTypeRole:
			value, ok := args[idx].(*discordgo.Role)
			if !ok {
				return "", fmt.Errorf("invalid argument type for component %s", baseID)
			}
			customID += "-" + value.ID
		}
	}

	return customID, nil
}

func (c *ButtonComponent) Build(args ...interface{}) (discordgo.MessageComponent, error) {
	id, err := getIDFromArgs(c.ID, c.Args, args)
	if err != nil {
		return nil, err
	}

	return &discordgo.Button{
		Label:    c.Label,
		Style:    c.Style,
		Disabled: c.Disabled,
		Emoji:    c.Emoji,
		CustomID: id,
	}, nil

}

func (c *SelectMenuComponent) Build(args ...interface{}) (discordgo.MessageComponent, error) {
	id, err := getIDFromArgs(c.ID, c.Args, args)
	if err != nil {
		return nil, err
	}

	return &discordgo.SelectMenu{
		CustomID:    id,
		Placeholder: c.Placeholder,
		MinValues:   c.MinValues,
		MaxValues:   c.MaxValues,
		Options:     c.Options,
		Disabled:    c.Disabled,
	}, nil
}
