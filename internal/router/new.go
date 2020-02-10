package router

import (
    "regexp"
    "github.com/bwmarrin/discordgo"
    rts "github.com/arturoguerra/d2arena/pkg/router"
)
var prefix = '-'

func New() *rts.Route {
    r :=  new(rts.Route)
    e := newHandler(r)
    r.EventHandler = e

    return r
}

func getargs(slices []string) ([]string) {
    startidx := 1
    args := make([]string, len(slices) - startidx)

    for i := startidx; i < len(slices); i++ {
        args[i - startidx] = slices[i]
    }

    return args
}

func newHandler(r *rts.Route) func(s *discordgo.Session, m *discordgo.MessageCreate) {
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

        restr := `\-([\w\d]+).*`
        re, err := regexp.Compile(restr)
        if err != nil {
            return
        }

        slice := re.FindStringSubmatch(m.Content)
        if len(slice) == 0 {
            return
        }

        name := slice[1]
        args := getargs(slice)

        if rt := r.Find(name); rt != nil {
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
