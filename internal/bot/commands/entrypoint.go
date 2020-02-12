package commands

import (
    "fmt"
    "github.com/arturoguerra/d2arena/pkg/router"
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/arturoguerra/d2arena/internal/logging"
)

var (
    log    = logging.New()
    grpcfg = config.LoadgRPC()
    grpcfaceit   = fmt.Sprintf("%s:%s", grpcfg.FaceitHost, grpcfg.FaceitPort)
    grpcprofiles = fmt.Sprintf("%s:%s", grpcfg.ProfilesHost, grpcfg.ProfilesPort)
)

func New(r *router.Route) {
    //r.On("ban", ban, true)
    //r.On("unban", unban, true)
    r.On("clear", clear, true)
    r.On("profile", profile, true)
}
