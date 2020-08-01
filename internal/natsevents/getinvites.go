package natsevents

import (
	"context"
	"fmt"

	"github.com/arturoguerra/d2arena/pkg/faceit"
	"google.golang.org/grpc"
)

func (h *handler) getInvite(hubid string) (string, error) {
	conn, err := grpc.Dial(h.Config.GRPC.Faceit, grpc.WithInsecure())
	if err != nil {
		return "", err
	}

	defer conn.Close()

	c := faceit.NewFaceitClient(conn)
	h.Logger.Infof("Fetching invites for %s\n", hubid)
	r, err := c.GetInvite(context.Background(), &faceit.InviteRequest{
		Hubid: hubid,
	})
	if err != nil {
		return "", err
	}

	link := fmt.Sprintf("%s/%s", r.GetBase(), r.GetCode())
	return link, nil
}
