package profiles

import (
	"context"

	pb "github.com/destinyarena/discord-bot/pkg/profiles"
)

func (p *profiles) Get(id string) (*Profile, error) {
	client, conn, err := p.conn()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	p.Logger.Info().Msgf("getting profile %s", id)

	r, err := client.GetProfile(context.Background(), &pb.IdRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &Profile{
		Discord:   r.GetDiscord(),
		Faceit:    r.GetFaceit(),
		Bungie:    r.GetBungie(),
		Banned:    r.GetBanned(),
		BanReason: r.GetBanreason(),
		IPHash:    r.GetIphash(),
	}, nil
}
