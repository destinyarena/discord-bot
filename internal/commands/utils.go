package commands

import (
	"errors"
	"regexp"

	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
	"github.com/labstack/gommon/log"
)

// GetUID returns UID i guess
func (bc *BaseCommand) GetUID(ctx *gommand.Context) (uid string, err error) {
	if len(ctx.Message.Mentions) != 0 {
		log.Info("Searching by Discord Mention")
		uid = ctx.Message.Mentions[0].ID.String()
	} else if len(ctx.Args) != 0 {
		id := ctx.Args[0].(string)
		log.Info(id)
		if m, _ := regexp.Match(`^\d+$`, []byte(id)); m {
			log.Info("Searching by Discord ID")
			uid = id
		} else if m, _ := regexp.Match(`^([A-f0-9\-])+$`, []byte(id)); m {
			log.Info("Searching by faceit GUID")
			uid = id
		} else {
			uid, err := bc.Faceit.GetIDByName(id)
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

func (bc *BaseCommand) StringArgs(args []interface{}) []string {
	sargs := make([]string, len(args))
	for i := 0; len(args) > i; i++ {
		sargs[i] = args[0].(string)
	}

	return sargs
}

type permFunc func(ctx *gommand.Context) (string, bool)

func inStringSlice(slice []string, item string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func (bc *BaseCommand) isOwner() permFunc {
	return func(ctx *gommand.Context) (string, bool) {
		if inStringSlice(bc.Config.Owners, ctx.Message.Author.ID.String()) {
			return "", true
		}

		return "Sorry you are not a bot owner", false
	}
}

func (bc *BaseCommand) isAdmin() permFunc {
	return func(ctx *gommand.Context) (string, bool) {
		guild := ctx.Session.Guild(disgord.ParseSnowflakeString(bc.Config.GuildID))
		member, err := guild.Member(ctx.Message.Author.ID).Get()
		if err != nil {
			return "Error getting member", false
		}

		for _, role := range member.Roles {
			for _, admrole := range bc.Config.AdminRoles {
				if role.String() == admrole {
					return "", true
				}
			}
		}

		return "Sorry you are not an admin", false
	}
}

func (bc *BaseCommand) isMod() permFunc {
	return func(ctx *gommand.Context) (string, bool) {
		guild := ctx.Session.Guild(disgord.ParseSnowflakeString(bc.Config.GuildID))
		member, err := guild.Member(ctx.Message.Author.ID).Get()
		if err != nil {
			return "Error getting member", false
		}

		for _, role := range member.Roles {
			for _, modrole := range bc.Config.ModRoles {
				if role.String() == modrole {
					return "", true
				}
			}
		}

		return "Sorry you are not an admin", false
	}
}
