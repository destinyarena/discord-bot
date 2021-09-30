package profiles

import (
	pb "github.com/destinyarena/bot/pkg/profiles"
	"google.golang.org/grpc"
)

func (p *profiles) conn() (pb.ProfilesClient, *grpc.ClientConn, error) {
	p.Logger.Info("GRPC Conn")
	conn, err := grpc.Dial(p.URL, grpc.WithInsecure())
	p.Logger.Info(conn)
	p.Logger.Info(err)
	if err != nil {
		p.Logger.Error(err)
		return nil, nil, err
	}

	client := pb.NewProfilesClient(conn)
	p.Logger.Info(client)

	return client, conn, nil

}
