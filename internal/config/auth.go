package config

import "os"

func LoadAuth() string {
    return os.Getenv("AUTH_TOKEN")
}
