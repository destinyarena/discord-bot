package commands

import (
	"context"
	"io"

	pb "github.com/arturoguerra/d2arena/pkg/faceit"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

// Hub is an individual hub
type Hub struct {
	Hubid  string
	Name   string
	GameID string
}

// GetHubs returns a list of hubs
func (c *Commands) GetHubs(guid string) ([]*Hub, error) {
	conn, err := grpc.Dial(c.Config.GRPC.Faceit, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	f := pb.NewFaceitClient(conn)
	stream, err := f.GetUserHubs(context.Background(), &pb.ProfileRequest{
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
