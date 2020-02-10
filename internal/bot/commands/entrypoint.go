package commands

import (
    "github.com/arturoguerra/d2arena/pkg/router"
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/arturoguerra/d2arena/internal/logging"
)

var (
    log = logging.New()
    grpcfg = config.LoadgRPC()
)

func New(r *router.Route) {
    //r.On("ban", ban, true)
    //r.On("unban", unban, true)
    r.On("clear", clear, true)
    r.On("profile", profile, true)
}
