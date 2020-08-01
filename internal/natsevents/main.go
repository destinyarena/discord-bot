package natsevents

import (
	"fmt"

	"github.com/arturoguerra/d2arena/internal/config"
	"github.com/arturoguerra/d2arena/internal/structs"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func (h *handler) register(s *discordgo.Session, id string) {
	if _, err := s.Guild(h.Config.Discord.GuildID); err != nil {
		h.Logger.Error(err)
		return
	}

	u, err := s.User(id)
	if err != nil {
		h.Logger.Error(err)
		return
	}

	channel, err := s.UserChannelCreate(id)
	if err != nil {
		h.Logger.Error(err)
		return
	}

	hubs := ""
	var roles []string

	for _, hub := range h.Config.Discord.Hubs {
		hubs += fmt.Sprintf("[%s](%s)\n", hub.Format, hub.HubID)
		roles = append(roles, hub.RoleID)
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Destiny Arena Faceit Invitation",
		Description: hubs,
	}

	logembed := &discordgo.MessageEmbed{
		Title:       "Hub invites Notification",
		Description: fmt.Sprintf("Sent hub invites to <@%s>(`%s#%s`)", id, u.Username, u.Discriminator),
	}

	if _, err := s.ChannelMessageSendEmbed(channel.ID, embed); err == nil {
		s.GuildMemberRoleAdd(h.Config.Discord.GuildID, id, h.Config.Discord.RegistrationRoleID)
	} else {
		logembed = &discordgo.MessageEmbed{
			Title:       "403: Forbidden",
			Description: fmt.Sprintf("Error sending hub channel to <@%s>(`%s#%s`) please contact them", id, u.Username, u.Discriminator),
		}
	}

	s.ChannelMessageSendEmbed(h.Config.Discord.LogsID, logembed)
}

func (h *handler) registration(dg *discordgo.Session, nchan *structs.NATS) {
	for i := range nchan.RecvRegistration {
		if i.Id != "" {
			h.Logger.Infof("Registering user: %s", i.Id)
			h.register(dg, i.Id)
		}
	}
}

type handler struct {
	Session *discordgo.Session
	Config  *config.Config
	Logger  *logrus.Logger
}

// New registeres
func New(dg *discordgo.Session, cfg *config.Config, logger *logrus.Logger, nchan *structs.NATS) {
	logger.Infoln("Registering NATS Events")
	h := &handler{
		Session: dg,
		Config:  cfg,
		Logger:  logger,
	}

	go h.registration(dg, nchan)
}
