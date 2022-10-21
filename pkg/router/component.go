package router

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type (
	ComponentHandlerFunc func(ctx *ComponentContext)

	ComponentContext struct {
		Context
		CustomID string
		Args     map[string]*ComponentArgument
	}

	ComponentRouter struct {
		components map[string]*Component
	}

	ComponentArgumentType int
	ComponentType         int
	ComponentArgument     struct {
		Name  string
		Value interface{}
		Type  ComponentArgumentType
	}

	Component struct {
		ID      string
		Type    ComponentType
		Args    []*ComponentArgument
		Handler ComponentHandlerFunc
	}
)

const (
	ComponentTypeButton ComponentType = iota
	ComponentTypeSelectMenu
)

const (
	ComponentArgumentTypeString ComponentArgumentType = iota
	ComponentArgumentTypeUser
	ComponentArgumentTypeChannel
	ComponentArgumentTypeRole
)

func NewComponentRouter(components []*Component) *ComponentRouter {
	router := &ComponentRouter{
		components: make(map[string]*Component),
	}

	return router
}

func (r *ComponentRouter) Register(components ...*Component) error {
	for _, component := range components {
		if _, ok := r.components[component.ID]; ok {
			return fmt.Errorf("component %s already registered", component.ID)
		}

		r.components[component.ID] = component
	}

	return nil
}

func (r *ComponentRouter) Get(id string) *Component {
	return r.components[id]
}

func getIDFromArgs(baseID string, cargs []*ComponentArgument, args []interface{}) (string, error) {
	if len(args) != len(cargs) {
		return baseID, errors.New("error: invalid number of component arguments")
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
