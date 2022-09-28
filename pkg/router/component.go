package router

import "github.com/bwmarrin/discordgo"

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
		Name    string
		Type    discordgo.ComponentType
		Args    []*ComponentArgument
		Handler ComponentHandlerFunc
	}
)

const (
	ComponentArgumentTypeString ComponentArgumentType = iota
	ComponentArgumentTypeUser
	ComponentArgumentTypeChannel
	ComponentArgumentTypeRole
)
