package structs

type (
	NATSConfig struct {
		URL string
	}

	NATSRegistration struct {
		Id string `json:"id"`
	}

	NATS struct {
		RecvRegistration chan *NATSRegistration
	}
)
