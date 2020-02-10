package commands

import (
    "github.com/arturoguerra/d2arena/pkg/router"
)

func profile(ctx *router.Context) {
    var searchType int
    var searchname string
    var _ = searchType
    var _ = searchname
    if len(ctx.Mentions) != 0 {
    } else if len(ctx.Args) != 0 {

    }
}
