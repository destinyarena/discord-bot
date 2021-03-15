package handlers

import (
	"github.com/andersfylling/disgord"
)

// OnMemberJoin handles the on_member_join discord event
func (h *handler) OnMemberJoin() func(disgord.Session, *disgord.GuildMemberAdd) {
	return func(s disgord.Session, gma *disgord.GuildMemberAdd) {
		profiles, err := h.Profiles.Get(gma.Member.User.ID.String())
		if err != nil {
			h.Logger.Error(err)
			return
		}

		if profiles.Banned == true {
			h.Logger.Infof("Looks like %s was banned", gma.Member.User.ID.String())
			return
		}

		h.Bot.Client.Guild(disgord.ParseSnowflakeString(h.Bot.Config.GuildID)).Member(gma.Member.User.ID).AddRole(disgord.ParseSnowflakeString(h.Bot.Config.RegistrationRoleID))

	}
}
