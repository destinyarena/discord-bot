package faceit

import (
	"context"

	pb "github.com/destinyarena/bot/pkg/faceit"
)

func (f *faceit) Ban(hubid, guid, reason string) error {
	client, conn, err := f.conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := client.Ban(context.Background(), &pb.BanRequest{
		Hubid:  hubid,
		Guid:   guid,
		Reason: reason,
	}); err != nil {
		return err
	}

	return nil
}
