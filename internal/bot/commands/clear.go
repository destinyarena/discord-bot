package commands

import (
    "fmt"
    "context"
    "google.golang.org/grpc"
    "github.com/arturoguerra/d2arena/internal/bot/utils"
    "github.com/arturoguerra/d2arena/pkg/router"
    "github.com/arturoguerra/d2arena/pkg/profiles"
)

func clear(ctx *router.Context) {
    err, uid := utils.GetUID(ctx)
    if err != nil{
        ctx.Reply(err.Error())
        return
    }

    addr := fmt.Sprintf("%s:%s", grpcfg.ProfilesHost, grpcfg.ProfilesPort)
    conn, err := grpc.Dial(addr, grpc.WithInsecure())
    if err != nil {
        log.Error(err)
        ctx.Reply("Error while connecting to ban systems")
        return
    }
    defer conn.Close()

    c := profiles.NewProfilesClient(conn)

    _, err = c.RemoveProfile(context.Background(), &profiles.IdRequest{
        Id: uid,
    })
    if err != nil {
        log.Error(err)
        ctx.Reply("Error while deleteting user profile")
        return
    }

    ctx.Reply("Deleted user profile")
}
