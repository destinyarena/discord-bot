package profiles

import (
	"context"

	pb "github.com/destinyarena/bot/pkg/profiles"
)

func (p *profiles) Remove(id string) error {
	client, conn, err := p.conn()
	if err != nil {
		return err
	}

	defer conn.Close()

	if _, err := client.RemoveProfile(context.Background(), &pb.IdRequest{
		Id: id,
	}); err != nil {
		return err
	}

	return nil
}
