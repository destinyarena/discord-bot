package commands

import (
	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
	"github.com/labstack/gommon/log"
)

var unbanhubs = []string{"f2c858d2-4a8e-42c0-830f-e97dda8da1ba"}

type unban struct {
	*BaseCommand
	gommand.CommandBasics
}

func (c *unban) Init() {
	c.Name = "unban"
	c.Description = "Unbans a member"
	c.PermissionValidators = []func(*gommand.Context) (string, bool){
		c.isOwner(),
		c.isAdmin(),
	}
}

func (c *unban) CommandFunction(ctx *gommand.Context) error {
	berr := false
	guild := ctx.Session.Guild(disgord.ParseSnowflakeString(c.Config.GuildID))
	if _, err := guild.Get(); err != nil {
		ctx.Reply("Error fetching Guild ID")
		return nil
	}

	uid, err := c.GetUID(ctx)
	if err != nil {
		ctx.Reply(err.Error())
		return err
	}

	profile, err := c.Profiles.Get(uid)
	if err != nil {
		ctx.Reply("Error fetching user profile")
		return err
	}

	if err = guild.UnbanUser(disgord.ParseSnowflakeString(profile.Discord), "True Pain"); err != nil {
		ctx.Reply("Error unbanning user from discord")
		berr = true
	}

	if err = c.Profiles.UnBan(uid); err != nil {
		ctx.Reply(err.Error())
		berr = true
	}

	for _, hub := range unbanhubs {
		if err := c.Faceit.UnBan(hub, profile.Faceit); err != nil {
			log.Error(err)
		}
	}

	if berr == false {
		ctx.Reply("Unbanned user successfully")
	}

	return nil
}
