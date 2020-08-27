package commands

import (
	"context"
	"errors"
	"regexp"
	"strings"

	faceit "github.com/arturoguerra/d2arena/pkg/faceit"
	"github.com/arturoguerra/d2arena/pkg/router"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

func (c *Commands) getFaceitGUIDByName(name string) (string, error) {
	conn, err := grpc.Dial(c.Config.GRPC.Faceit, grpc.WithInsecure())
	if err != nil {
		return "", err
	}

	defer conn.Close()

	f := faceit.NewFaceitClient(conn)
	log.Info("Fetching faceit profile")
	r, err := f.GetProfileByName(context.Background(), &faceit.ProfileNameRequest{
		Name: name,
	})

	if err != nil {
		return "", err
	}

	return r.GetGuid(), nil
}

// GetUID returns UID i guess
func (c *Commands) GetUID(ctx *router.Context) (uid string, err error) {
	if len(ctx.Mentions) != 0 {
		log.Info("Searching by Discord Mention")
		uid = ctx.Mentions[0].ID
	} else if len(ctx.Args) != 0 {
		id := strings.Join(ctx.Args[0:], " ")
		log.Info(id)
		if m, _ := regexp.Match(`^\d+$`, []byte(id)); m {
			log.Info("Searching by Discord ID")
			uid = id
		} else if m, _ := regexp.Match(`^([A-f0-9\-])+$`, []byte(id)); m {
			log.Info("Searching by faceit GUID")
			uid = id
		} else {
			uid, err := c.getFaceitGUIDByName(id)
			if err != nil {
				return "", err
			}

			return uid, nil
		}
	} else {
		err = errors.New("Sorry but you must provide a way for us to find the user")
	}

	return uid, err
}
