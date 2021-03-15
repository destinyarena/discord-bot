package faceit

import (
	pb "github.com/destinyarena/bot/pkg/faceit"
	"google.golang.org/grpc"
)

func (f *faceit) conn() (pb.FaceitClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(f.URL, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewFaceitClient(conn)

	return client, conn, nil
}
