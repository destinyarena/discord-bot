package router

import (
    "fmt"
    "regexp"
    "strings"
    "github.com/bwmarrin/discordgo"
    "github.com/arturoguerra/d2arena/internal/config"
    rts "github.com/arturoguerra/d2arena/pkg/router"
    "github.com/arturoguerra/d2arena/internal/logging"
)
var prefix = "!"
var log = logging.New()

func getArgs(content string, command string) []string {
    trimed := strings.TrimLeft(content, fmt.Sprintf("%s%s", prefix, command))
    return strings.Fields(trimed)
}

func New() *rts.Route {
    r :=  new(rts.Route)
    r.EventHandler = newHandler(r)

    return r
}

func isAllowed(m *discordgo.Member) bool {
    dcfg := config.LoadDiscord()
    for _, role := range m.Roles {
        for _, arole := range dcfg.StaffRoles {
            if role == arole {
                return true
            }
        }
    }

    return false
}

func isOwner(uid string) bool {
    dcfg := config.LoadDiscord()
    for _, user := range dcfg.Owners {
        if user == uid {
            return true
        }
    }

    return false
}

func newHandler(r *rts.Route) func(*discordgo.Session, *discordgo.MessageCreate) {
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
                log.Info("Command set to admin only, checking perms..")
                if m.Member != nil && isAllowed(m.Member) {
                    log.Info("User is set as admin")
                    allowed = true
                }

                if isOwner(m.Author.ID) {
                    log.Info("User is an owner")
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
