package router

import (
    "github.com/bwmarrin/discordgo"
    "strings"
    "errors"
    "regexp"
)


type (
    Context struct {
        *discordgo.Message
        Channel *discordgo.Channel
        Guild *discordgo.Guild
        Session *discordgo.Session
    }

    HandlerFunc func(ctx *Context)

    Route struct {
        Name string
        Handler HandlerFunc
    }

    Router struct {
       Routes []*Route
       Session *discordgo.Session
       Prefix string
    }
)

var ErrRouteAlreadyRegistered = errors.New("Route has already been registered")

// Route Methods
func (rt *Route) Match(name string) bool {
    if strings.ToLower(name) == strings.ToLower(rt.Name) {
        return true
    }

    return false
}

// Router Methods
func New(s *discordgo.Session, c *Config) *Router {
    return &Router{
        Session: s,
        Prefix: c.Prefix,
    }
}

func (r *Router) Handler(m *discordgo.MessageCreate) error {
    if r.Session.State.User.ID == m.Author.ID {
        return nil
    }

    c, err := r.Session.State.Channel(m.ChannelID)
    if err != nil {
        return nil
    }

    g, err := r.Session.State.Guild(c.GuildID)
    if err != nil {
        return nil
    }

    restr := `\` + r.Prefix + `([\w\d]+).*`
    re, err := regexp.Compile(restr)

    if err != nil {
        return err
    }

    contentslice := re.FindStringSubmatch(m.Content)

    if len(contentslice) == 0 {
        return nil
    }

    name := contentslice[1]

    if rt := r.Find(name); rt != nil {
        ctx := &Context{
            m.Message,
            c,
            g,
            r.Session,
        }

        rt.Handler(ctx)
    }

    return nil
}

func (r *Router) Find(name string) *Route {
    for _, rt := range r.Routes {
        if rt.Match(name) {
            return rt
        }
    }

    return nil
}

func (r *Router) On(name string, handler HandlerFunc) error {
    route := &Route{
        Name: name,
        Handler: handler,
    }

    if rt := r.Find(route.Name); rt != nil {
        return ErrRouteAlreadyRegistered
    }

    r.Routes = append(r.Routes, route)
    return nil
}
