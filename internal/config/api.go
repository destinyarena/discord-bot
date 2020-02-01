package config

import (
    "os"
    "github.com/arturoguerra/d2arena/internal/structs"
)

func LoadAPIConfig() *structs.API {
    return &structs.API{
        BaseURL: os.Getenv("API_BASE_URL"),
    }
}
