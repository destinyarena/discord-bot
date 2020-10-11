package router

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/bot/internal/config"
	rts "github.com/destinyarena/bot/pkg/router"
	"github.com/sirupsen/logrus"
)

var prefix = "!"

func getArgs(content string, command string) []string {
	trimed := strings.TrimLeft(content, fmt.Sprintf("%s%s", prefix, command))
	return strings.Fields(trimed)
}

// Route is a sub structure for rts.Route that includes a config
type Route struct {
	*rts.Route
	Logger *logrus.Logger
	Config *config.Config
}

// New returns a new instance of a discord command router
func New(cfg *config.Config, log *logrus.Logger) *Route {
	r := &Route{Route: new(rts.Route), Config: cfg, Logger: log}
	r.Route.EventHandler = r.newHandler()

	return r
}

func (r *Route) isAllowed(m *discordgo.Member) bool {
	for _, role := range m.Roles {
		for _, arole := range r.Config.Discord.StaffRoles {
			if role == arole {
				return true
			}
		}
	}

	return false
}

func (r *Route) isOwner(uid string) bool {
	for _, user := range r.Config.Discord.Owners {
		if user == uid {
			return true
		}
	}

	return false
}

func (r *Route) newHandler() func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if s.State.User.ID == m.Author.ID {
			return
		}

		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			return
		}

		g, err := s.State.Guild(c.GuildID)
		if err != nil {
			return
		}

		restr := `\` + prefix + `([\w\d]+).*`
		re, err := regexp.Compile(restr)
		if err != nil {
			return
		}

		slice := re.FindStringSubmatch(m.Content)
		if len(slice) == 0 {
			return
		}

		name := slice[1]
		args := getArgs(m.Content, name)

		if rt := r.Find(name); rt != nil {
			allowed := true
			if rt.Admin {
				allowed = false
				r.Logger.Info("Command set to admin only, checking perms..")
				if m.Member != nil && r.isAllowed(m.Member) {
					r.Logger.Info("User is set as admin")
					allowed = true
				}

				if r.isOwner(m.Author.ID) {
					r.Logger.Info("User is an owner")
					allowed = true
				}
			}

			if allowed {
				ctx := &rts.Context{
					Message: m.Message,
					Channel: c,
					Guild:   g,
					Session: s,
					Args:    args,
					Command: name,
				}

				rt.Handler(ctx)
			}

		}
	}
}
