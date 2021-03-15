package faceit

import (
	"context"

	pb "github.com/destinyarena/bot/pkg/faceit"
)

func (f *faceit) GetProfileByID(guid string) (*Profile, error) {
	client, conn, err := f.conn()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	r, err := client.GetProfileByID(context.Background(), &pb.ProfileRequest{
		Guid: guid,
	})
	if err != nil {
		return nil, err
	}

	return &Profile{
		GUID:     r.GetGuid(),
		Username: r.GetUsername(),
		Level:    int(r.GetSkilllvl()),
	}, nil
}
