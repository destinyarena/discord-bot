package config

import "github.com/joeshaw/envdecode"

func loadEnv(i interface{}) error {
	return envdecode.Decode(i)
}
