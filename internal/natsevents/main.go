package natsevents

import (
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/destinyarena/bot/internal/bot"
	"github.com/destinyarena/bot/internal/structs"
	"github.com/sirupsen/logrus"
)

func (h *handler) register(id string) {

	user := h.Bot.Client.User(disgord.ParseSnowflakeString(id))
	channel, err := user.CreateDM()
	if err != nil {
		h.Logger.Error(err)
		return
	}

	u, err := user.Get()
	if err != nil {
		h.Logger.Error(err)
		return
	}

	hubs := ""
	var roles []string

	for _, hub := range h.Bot.Config.Hubs {
		link, err := h.getInvite(hub.HubID)
		if err != nil {
			h.Logger.Error(err)
		} else {
			hubs += fmt.Sprintf("[%s](%s) \n", hub.Format, link)
			roles = append(roles, hub.RoleID)
		}
	}

	embed := &disgord.Embed{
		Description: hubs,
	}

	logembed := &disgord.Embed{
		Description: fmt.Sprintf("Sent hub invites to <@%s>(`%s#%s`)", id, u.Username, u.Discriminator),
	}

	h.Bot.Client.Guild(disgord.ParseSnowflakeString(h.Bot.Config.GuildID)).Member(u.ID).AddRole(disgord.ParseSnowflakeString(h.Bot.Config.RegistrationRoleID))

	if _, err := h.Bot.Client.SendMsg(channel.ID, embed); err != nil {
		logembed = &disgord.Embed{
			Title:       "403: Forbidden",
			Description: fmt.Sprintf("Error sending hub channel to <@%s>(`%s#%s`) please contact them", id, u.Username, u.Discriminator),
		}
	}

	h.Bot.Client.SendMsg(disgord.ParseSnowflakeString(h.Bot.Config.LogsID), logembed)
}

func (h *handler) registration(nchan *structs.NATS) {
	for i := range nchan.RecvRegistration {
		if i.Id != "" {
			h.Logger.Infof("Registering user: %s", i.Id)
			h.register(i.Id)
		}
	}
}

type handler struct {
	Bot    *bot.Bot
	Logger *logrus.Logger
}

// New registeres
func New(bot *bot.Bot, logger *logrus.Logger, nchan *structs.NATS) {
	logger.Infoln("Registering NATS Events")
	h := &handler{
		Bot:    bot,
		Logger: logger,
	}

	go h.registration(nchan)
}
