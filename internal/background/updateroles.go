package background

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/bot/internal/config"
)

func checkroles(roles []string, roleid string) bool {
	for _, role := range roles {
		if role == roleid {
			return false
		}
	}

	return true
}

func filter(members []*discordgo.Member, rid string) []*discordgo.Member {
	var mlist []*discordgo.Member
	for _, member := range members {
		createdAt, _ := discordgo.SnowflakeTimestamp(member.User.ID)
		loc, _ := time.LoadLocation("UTC")
		now := time.Now().In(loc)
		diff := now.Sub(createdAt)

		days := int(diff.Hours()) / 24

		if days >= 7 && checkroles(member.Roles, rid) {
			mlist = append(mlist, member)
		}
	}

	return mlist
}

func (bh *background) UpdateRoles() {
	for true {
		guild, err := bh.Session.Guild(bh.Config.Discord.GuildID)
		if err == nil {
			members := filter(guild.Members, bh.Config.Discord.JoinRoleID)
			for _, member := range members {
				bh.Session.GuildMemberRoleAdd(guild.ID, member.User.ID, bh.Config.Discord.JoinRoleID)
			}
		}

		time.Sleep(120 * time.Second)
	}
}

type (
	background struct {
		Session *discordgo.Session
		Config  *config.Config
	}

	// Background handlers
	Background interface {
		UpdateRoles()
	}
)

// New returns a new background handler
func New(s *discordgo.Session, config *config.Config) Background {
	return &background{
		Session: s,
		Config:  config,
	}
}
