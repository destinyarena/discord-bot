package commands

import (
	"context"

	faceit "github.com/destinyarena/bot/pkg/faceit"
	profiles "github.com/destinyarena/bot/pkg/profiles"
	"github.com/destinyarena/bot/pkg/router"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

var unbanhubs = []string{"f2c858d2-4a8e-42c0-830f-e97dda8da1ba"}

func (c *Commands) unban(ctx *router.Context) {
	guild, err := ctx.Session.Guild(c.Config.Discord.GuildID)
	if err != nil {
		ctx.Reply("Error fetching Guild ID")
		return
	}

	uid, err := c.GetUID(ctx)
	if err != nil {
		ctx.Reply(err.Error())
	}

	profile, err := c.getProfile(uid)
	if err != nil {
		ctx.Reply("Error fetching user profile")
	}

	if err = ctx.Session.GuildBanDelete(guild.ID, profile.Discord); err != nil {
		ctx.Reply("Error unbanning user from discord")
		return
	}

	pconn, err := grpc.Dial(c.Config.GRPC.Profile, grpc.WithInsecure())
	if err != nil {
		ctx.Reply("Error conneting to internal profiles system")
		return
	}

	defer pconn.Close()

	pClient := profiles.NewProfilesClient(pconn)

	_, err = pClient.Unban(context.Background(), &profiles.IdRequest{
		Id: profile.Discord,
	})

	fconn, err := grpc.Dial(c.Config.GRPC.Faceit, grpc.WithInsecure())
	if err != nil {
		ctx.Reply("Error connecting to internal faceit system")
		return
	}

	defer fconn.Close()

	fClient := faceit.NewFaceitClient(fconn)

	for _, hub := range unbanhubs {
		if _, err := fClient.Unban(context.Background(), &faceit.UnbanRequest{
			Hubid: hub,
			Guid:  profile.Faceit,
		}); err != nil {
			log.Error(err)
		}
	}
}
