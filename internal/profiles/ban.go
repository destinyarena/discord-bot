package profiles

import (
	"context"

	pb "github.com/destinyarena/discord-bot/pkg/profiles"
)

func (p *profiles) Ban(id, reason string) error {
	client, conn, err := p.conn()
	if err != nil {
		return err
	}

	defer conn.Close()

	if _, err := client.Ban(context.Background(), &pb.BanRequest{
		Id:     id,
		Reason: reason,
	}); err != nil {
		return err
	}

	return nil
}
