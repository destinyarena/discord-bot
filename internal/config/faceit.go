package config

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "os"
)

func LoadFaceit() *structs.Faceit {
    utoken := os.Getenv("FACEIT_USER_TOKEN")
    apitoken := os.Getenv("FACEIT_API_TOKEN")

    return &structs.Faceit{
        utoken,
        apitoken,
    }
}
