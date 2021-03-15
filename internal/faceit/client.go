package faceit

import "github.com/sirupsen/logrus"

type (
	Hub struct {
		HubID  string
		Name   string
		GameID string
	}

	Profile struct {
		GUID     string
		Username string
		Level    int
	}

	// Faceit : exports client functions
	Faceit interface {
		Ban(hubid, guid, reason string) error
		UnBan(hubid, guid string) error
		GetUserHubs(guid string) ([]*Hub, error)
		GetProfileByID(guid string) (*Profile, error)
		GetIDByName(name string) (string, error)
	}

	// internal client struct
	faceit struct {
		*Config
		Logger *logrus.Logger
	}
)

func New(logger *logrus.Logger, config *Config) (Faceit, error) {
	f := &faceit{
		Config: config,
		Logger: logger,
	}

	return f, nil
}
