package config

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "os"
)

func LoadDiscord() *structs.Discord {
    prefix := os.Getenv("PREFIX")
    token := os.Getenv("DTOKEN")
    joinrole := os.Getenv("DJOINROLEID")

    return &structs.Discord{
        prefix,
        token,
        joinrole,
    }
}
