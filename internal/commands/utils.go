package commands

import (
	"regexp"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/snowflake/v5"
	"github.com/auttaja/gommand"
	"github.com/labstack/gommon/log"
)

// GetUID returns UID i guess
func (bc *BaseCommand) GetUID(ctx *gommand.Context) (uid string, err error) {
	if len(ctx.Message.Mentions) != 0 {
		log.Info("Searching by Discord Mention")
		uid = ctx.Message.Mentions[0].ID.String()
	} else {
		switch v := ctx.Args[0].(type) {
		case *disgord.User:
			return v.ID.String(), nil
		case string:
			if m, _ := regexp.Match(`^([A-f0-9\-])+$`, []byte(v)); m {
				log.Info("Searching by faceit GUID")
				return v, nil
			} else {
				return bc.Faceit.GetIDByName(v)
			}
		}
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

//var permFunc func(ctx *gommand.Context) (string, bool)

func inStringSlice(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func hasRole(slice []snowflake.Snowflake, roles []string) bool {
	for _, v := range slice {
		for _, rid := range roles {
			if v.String() == rid {
				return true
			}
		}
	}

	return false
}

func (bc *BaseCommand) permLvl(ctx *gommand.Context) int {
	if inStringSlice(bc.Config.Owners, ctx.Message.Author.ID.String()) {
		return 10
	}

	member, err := ctx.Session.Guild(disgord.ParseSnowflakeString(bc.Config.GuildID)).Member(ctx.Message.Author.ID).Get()
	if err != nil {
		return 0
	}

	// Admin
	if hasRole(member.Roles, bc.Config.AdminRoles) {
		return 10
	}

	// Mod
	if hasRole(member.Roles, bc.Config.ModRoles) {
		return 5
	}

	return 0
}

func (bc *BaseCommand) isAllowed(permlvl int) []func(*gommand.Context) (string, bool) {
	return []func(*gommand.Context) (string, bool){
		func(ctx *gommand.Context) (string, bool) {
			if bc.permLvl(ctx) >= permlvl {
				return "", true
			}

			return "Sorry but you don't have permission to run this command", false
		},
	}

}
