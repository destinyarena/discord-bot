package commands

import (
	"errors"
	"regexp"
	"strings"

	"github.com/arturoguerra/d2arena/pkg/router"
	"github.com/labstack/gommon/log"
)

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
			err = errors.New("Fetchng profiles with a faceit name is unavailable at the moment")
		}
	} else {
		err = errors.New("Sorry but you must provide a way for us to find the user")
	}

	return uid, err
}
