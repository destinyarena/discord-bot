package commands

import (
	"github.com/arturoguerra/d2arena/internal/router"
)

// Commands holds all commands with access to parent router
type Commands struct {
	*router.Route
}

// New returns a new command handler
func New(r *router.Route) {
	c := &Commands{Route: r}
	r.On("ban", c.ban, true)
	//r.On("unban", unban, true)
	r.On("clear", c.clear, true)
	r.On("profile", c.profile, true)
}
