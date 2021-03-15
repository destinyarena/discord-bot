package faceit

import (
	"context"

	pb "github.com/destinyarena/bot/pkg/faceit"
)

func (f *faceit) UnBan(hubid, guid string) error {
	client, conn, err := f.conn()
	if err != nil {
		return err
	}

	defer conn.Close()

	if _, err := client.Unban(context.Background(), &pb.UnbanRequest{
		Hubid: hubid,
		Guid:  guid,
	}); err != nil {
		return err
	}

	return nil
}
