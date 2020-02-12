package utils

import (
    "io"
    "context"
    "google.golang.org/grpc"
    pb "github.com/arturoguerra/d2arena/pkg/faceit"
)

type Hub struct {
    Hubid  string
    Name   string
    GameID string
}

func GetHubs(guid string) ([]*Hub, error) {
    conn, err := grpc.Dial(grpcfaceit, grpc.WithInsecure())
    if err != nil {
        return nil, err
    }

    defer conn.Close()

    c := pb.NewFaceitClient(conn)
    stream, err := c.GetUserHubs(context.Background(), &pb.ProfileRequest{
        Guid: guid,
    })
    if err != nil {
        return nil, err
    }

    log.Infof("Getting hubs for: %s", guid)
    hubs := make([]*Hub, 0)

    for {
        hub, err := stream.Recv()
        if err == io.EOF {
            log.Infof("End of hubs stream for: %s", guid)
            break
        }

        if err != nil {
            log.Error("Error: %s", err.Error())
            return nil, err
        }

        log.Infof("Hub ID: %s  Name: %s GameID: %s", hub.GetHubid(), hub.GetName(), hub.GetGameid())

        hubs = append(hubs, &Hub{
            Hubid:  hub.GetHubid(),
            Name:   hub.GetName(),
            GameID: hub.GetGameid(),
        })
    }

    return hubs, nil
}
