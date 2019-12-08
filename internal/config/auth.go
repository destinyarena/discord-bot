package config

import "os"

func LoadAuth() string {
    return os.Getenv("AUTHTOKEN")
}
