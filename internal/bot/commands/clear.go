package commands

import (
	"context"

	"github.com/arturoguerra/d2arena/pkg/profiles"
	"github.com/arturoguerra/d2arena/pkg/router"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

func (c *Commands) clear(ctx *router.Context) {
	uid, err := c.GetUID(ctx)
	if err != nil {
		ctx.Reply(err.Error())
		return
	}

	conn, err := grpc.Dial(c.Config.GRPC.Profile, grpc.WithInsecure())
	if err != nil {
		c.Logger.Error(err)
		ctx.Reply("Error while connecting to ban systems")
		return
	}
	defer conn.Close()

	p := profiles.NewProfilesClient(conn)

	_, err = p.RemoveProfile(context.Background(), &profiles.IdRequest{
		Id: uid,
	})
	if err != nil {
		log.Error(err)
		ctx.Reply("Error while deleteting user profile")
		return
	}

	ctx.Reply("Deleted user profile")
}
