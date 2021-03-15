package faceit

import (
	"context"

	pb "github.com/destinyarena/bot/pkg/faceit"
)

func (f *faceit) GetIDByName(name string) (string, error) {
	client, conn, err := f.conn()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	r, err := client.GetProfileByName(context.Background(), &pb.ProfileNameRequest{
		Name: name,
	})
	if err != nil {
		return "", err
	}

	return r.GetGuid(), nil
}
