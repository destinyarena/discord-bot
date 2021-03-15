package profiles

import "github.com/sirupsen/logrus"

type (
	Profile struct {
		Discord   string
		Faceit    string
		Bungie    string
		BanReason string
		IPHash    string
		Banned    bool
	}

	Profiles interface {
		UnBan(id string) error
		Ban(id, reason string) error
		Get(id string) (*Profile, error)
		Remove(id string) error
	}

	profiles struct {
		*Config
		Logger *logrus.Logger
	}
)

func New(logger *logrus.Logger, config *Config) (Profiles, error) {
	p := &profiles{
		config,
		logger,
	}

	return p, nil
}
