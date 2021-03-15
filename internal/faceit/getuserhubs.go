package faceit

import (
	"context"
	"io"

	pb "github.com/destinyarena/bot/pkg/faceit"
)

func (f *faceit) GetUserHubs(guid string) ([]*Hub, error) {
	client, conn, err := f.conn()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	stream, err := client.GetUserHubs(context.Background(), &pb.ProfileRequest{
		Guid: guid,
	})
	if err != nil {
		return nil, err
	}

	f.Logger.Infof("Getting hubs for: %s", guid)
	hubs := make([]*Hub, 0)

	for {
		hub, err := stream.Recv()
		if err == io.EOF {
			f.Logger.Infof("End of hubs stream for: %s", guid)
			break
		}

		if err != nil {
			f.Logger.Error("Error: %s", err.Error())
			return nil, err
		}

		f.Logger.Infof("Hub ID: %s  Name: %s GameID: %s", hub.GetHubid(), hub.GetName(), hub.GetGameid())

		hubs = append(hubs, &Hub{
			HubID:  hub.GetHubid(),
			Name:   hub.GetName(),
			GameID: hub.GetGameid(),
		})
	}

	return hubs, nil

}
