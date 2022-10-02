package router

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type (
	ModalHandlerFunc func(ctx *ModalContext)

	ModalContext struct {
		Context
		CustomID string
		Args     map[string]*ModalArgument
	}

	ModalArgumentType int

	ModalArgument struct {
		Name  string
		Value interface{}
		Type  ModalArgumentType
	}

	Modal struct {
		ID          string
		Label       string
		Style       discordgo.TextInputStyle
		Placeholder string
		Value       string
		Required    bool
		MinLength   int
		MaxLength   int
		Handler     ModalHandlerFunc
		Args        []*ModalArgument
	}
)

const (
	ModalArgumentTypeString ModalArgumentType = iota
	ModalArgumentTypeUser
	ModalArgumentTypeChannel
	ModalArgumentTypeRole
)

func (m *Modal) BuildComponent(args ...interface{}) (discordgo.MessageComponent, error) {
	if len(args) != len(m.Args) {
		return nil, fmt.Errorf("invalid number of arguments for modal %s", m.ID)
	}

	customID := m.ID
	for idx, arg := range m.Args {
		switch arg.Type {
		case ModalArgumentTypeString:
			value, ok := args[idx].(string)
			if !ok {
				return nil, fmt.Errorf("invalid argument type for modal %s", m.ID)
			}
			customID += fmt.Sprintf(":%s", value)
		case ModalArgumentTypeUser:
			value, ok := args[idx].(*discordgo.User)
			if !ok {
				return nil, fmt.Errorf("invalid argument type for modal %s", m.ID)
			}
			customID += fmt.Sprintf(":%s", value.ID)
		case ModalArgumentTypeChannel:
			value, ok := args[idx].(*discordgo.Channel)
			if !ok {
				return nil, fmt.Errorf("invalid argument type for modal %s", m.ID)
			}
			customID += fmt.Sprintf(":%s", value.ID)
		case ModalArgumentTypeRole:
			value, ok := args[idx].(*discordgo.Role)
			if !ok {
				return nil, fmt.Errorf("invalid argument type for modal %s", m.ID)
			}
			customID += fmt.Sprintf(":%s", value.ID)
		}
	}

	return &discordgo.TextInput{
		CustomID:    customID,
		Label:       m.Label,
		Style:       m.Style,
		Placeholder: m.Placeholder,
		Value:       m.Value,
		Required:    m.Required,
		MinLength:   m.MinLength,
		MaxLength:   m.MaxLength,
	}, nil

}
