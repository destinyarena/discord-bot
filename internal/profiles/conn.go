package profiles

import (
	pb "github.com/destinyarena/bot/pkg/profiles"
	"google.golang.org/grpc"
)

func (p *profiles) conn() (pb.ProfilesClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(p.URL, grpc.WithInsecure())
	if err != nil {
		p.Logger.Error(err)
		return nil, nil, err
	}

	client := pb.NewProfilesClient(conn)

	return client, conn, nil

}
