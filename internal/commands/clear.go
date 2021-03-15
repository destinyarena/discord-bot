package commands

import (
	"github.com/auttaja/gommand"
	"github.com/labstack/gommon/log"
)

type clear struct {
	gommand.CommandBasics
	BaseCommand
}

func (c *clear) Init() {
	c.Name = "clear"
	c.Description = "Clears a user profile"
	c.PermissionValidators = []func(*gommand.Context) (string, bool){
		c.isOwner(),
		c.isAdmin(),
	}
}

func (c *clear) CommandFunction(ctx *gommand.Context) error {
	uid, err := c.GetUID(ctx)
	if err != nil {
		ctx.Reply(err.Error())
		return err
	}

	if err = c.Profiles.Remove(uid); err != nil {
		log.Error(err)
		ctx.Reply("Error while deleteting user profile")
		return err
	}

	ctx.Reply("Deleted user profile")
	return nil
}
