package handlers

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/bot/pkg/profiles"
	"google.golang.org/grpc"
)

// OnMemberJoin handles the on_member_join discord event
func (h *handler) OnMemberJoin(s *discordgo.Session, member *discordgo.GuildMemberAdd) {
	conn, err := grpc.Dial(h.Config.GRPC.Profile, grpc.WithInsecure())
	if err != nil {
		h.Logger.Error(err.Error())
		return
	}

	defer conn.Close()

	p := profiles.NewProfilesClient(conn)
	r, err := p.GetProfile(context.Background(), &profiles.IdRequest{
		Id: member.User.ID,
	})

	if err != nil {
		h.Logger.Error(err.Error())
		return
	}

	if r.GetBanned() == true {
		h.Logger.Infof("Looks like %s was banned", member.User.Username)
		return
	}

	s.GuildMemberRoleAdd(h.Config.Discord.GuildID, member.User.ID, h.Config.Discord.RegistrationRoleID)
}
