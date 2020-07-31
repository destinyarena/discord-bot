package handlers

import (
	"context"

	faceit "github.com/arturoguerra/d2arena/pkg/faceit"
	profiles "github.com/arturoguerra/d2arena/pkg/profiles"
	"google.golang.org/grpc"

	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Profile does stuff
type Profile struct {
	Discord   string `json:"discord"   validate:"required"`
	Bungie    string `json:"bungie"    validate:"required"`
	Faceit    string `json:"faceit"    validate:"required"`
	Faceitlvl int    `json:"faceitlvl" validate:"required"`
}

func checkRole(roles []string, role string) bool {
	for _, r := range roles {
		if r == role {
			return false
		}
	}

	return true
}

func getMember(g *discordgo.Guild, uid string) (*discordgo.Member, error) {
	for _, member := range g.Members {
		if member.User.ID == uid {
			return member, nil
		}
	}

	return nil, errors.New("Not found")
}

func (h *handler) getInvite(hubid string) (string, error) {
	conn, err := grpc.Dial(h.Config.GRPC.Faceit, grpc.WithInsecure())
	if err != nil {
		h.Logger.Error(err)
		return "", err
	}
	defer conn.Close()

	c := faceit.NewFaceitClient(conn)
	h.Logger.Infoln("Fetching Invites")
	r, err := c.GetInvite(context.Background(), &faceit.InviteRequest{
		Hubid: hubid,
	})
	if err != nil {
		h.Logger.Error(err)
		return "", err
	}

	link := fmt.Sprintf("%s/%s", r.GetBase(), r.GetCode())
	return link, nil
}

func (h *handler) fetchProfile(id string) (*Profile, error) {
	conn, err := grpc.Dial(h.Config.GRPC.Profile, grpc.WithInsecure())
	if err != nil {
		h.Logger.Error(err)
		return nil, err
	}

	defer conn.Close()

	c := profiles.NewProfilesClient(conn)
	h.Logger.Infoln("Fetching database profile")
	r, err := c.GetProfile(context.Background(), &profiles.IdRequest{
		Id: id,
	})

	if err != nil {
		h.Logger.Error(err)
		return nil, err
	}

	conn, err = grpc.Dial(h.Config.GRPC.Faceit, grpc.WithInsecure())
	if err != nil {
		h.Logger.Error(err)
		return nil, err
	}

	defer conn.Close()
	h.Logger.Infoln("Fetching faceit skill level")

	f := faceit.NewFaceitClient(conn)
	rf, err := f.GetProfile(context.Background(), &faceit.ProfileRequest{
		Guid: r.GetFaceit(),
	})
	if err != nil {
		h.Logger.Error(err)
		return nil, err
	}

	return &Profile{
		Discord:   r.GetDiscord(),
		Bungie:    r.GetBungie(),
		Faceit:    r.GetFaceit(),
		Faceitlvl: int(rf.GetSkilllvl()),
	}, nil
}

func (h *handler) invites(s *discordgo.Session, mr *discordgo.MessageReactionAdd) {
	title := "Destiny Arena Faceit Invitation"

	guild, err := s.Guild(h.Config.Discord.GuildID)
	if err != nil {
		return
	}

	u, err := s.User(mr.UserID)
	if err != nil {
		return
	}

	channel, err := s.UserChannelCreate(mr.UserID)
	if err != nil {
		return
	}

	member, err := getMember(guild, mr.UserID)
	if err != nil {
		return
	}

	var hubs string
	send := false
	var roles []string
	var profile *Profile

	for _, hub := range h.Config.Discord.Hubs {
		if hub.EmojiID == mr.Emoji.APIName() && hub.MessageID == mr.MessageID {
			if profile == nil {
				profile, err = h.fetchProfile(mr.UserID)
				if err != nil {
					h.Logger.Errorln(err)
					embed := &discordgo.MessageEmbed{
						Title:       title,
						Description: "Looks like you haven't Registered yet, please do that before requesting invites.",
					}

					s.ChannelMessageSendEmbed(channel.ID, embed)
					return
				}
			}

			if checkRole(member.Roles, hub.RoleID) {
				h.Logger.Infoln("-----------")
				h.Logger.Infoln(profile.Faceitlvl)
				h.Logger.Infoln(hub.SkillLvl)
				if profile.Faceitlvl >= hub.SkillLvl {
					h.Logger.Infoln("Getting invite..")
					link, err := h.getInvite(hub.HubID)
					if err != nil || link == "" {
						h.Logger.Error(err)
					} else {
						roles = append(roles, hub.RoleID)
						hubs += fmt.Sprintf("[%s](%s)\n", hub.Format, link)
						send = true
					}
				}
			}
		}
	}

	if send {
		message := hubs
		embed := &discordgo.MessageEmbed{
			Title:       title,
			Description: message,
		}

		if _, err := s.ChannelMessageSendEmbed(channel.ID, embed); err == nil {
			for _, role := range roles {
				s.GuildMemberRoleAdd(h.Config.Discord.GuildID, mr.UserID, role)
			}
			s.ChannelMessageSendEmbed(h.Config.Discord.LogsID, &discordgo.MessageEmbed{
				Title:       "Notification",
				Description: fmt.Sprintf("User <@%s>(`%s#%s`) requested invites", mr.UserID, u.Username, u.Discriminator),
			})
		} else {
			embed := &discordgo.MessageEmbed{
				Title:       "403: Forbidden",
				Description: fmt.Sprintf("Error sending invites to <@%s>(`%s#%s`)", mr.UserID, u.Username, u.Discriminator),
			}

			s.ChannelMessageSendEmbed(h.Config.Discord.LogsID, embed)
		}
	}
}

func (h *handler) OnMessageReactionAdd(s *discordgo.Session, mr *discordgo.MessageReactionAdd) {
	h.invites(s, mr)
	h.reactionroles(s, mr.MessageReaction, true)
}
