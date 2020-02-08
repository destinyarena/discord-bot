package config

import (
    "os"
    "github.com/arturoguerra/d2arena/internal/structs"
)

func LoadNATSConfig() *structs.NATSConfig {
    return &structs.NATSConfig{
        URL: os.Getenv("NATS_URL"),
    }
}
