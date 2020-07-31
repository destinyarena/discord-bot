package config

// Auth loads the auth config
type Auth struct {
	AuthToken string `env:"AUTH_TOKEN,required"`
}
