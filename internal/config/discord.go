package config

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "os"
)

func LoadDiscord() *structs.Discord {
    prefix := os.Getenv("PREFIX")
    token := os.Getenv("DTOKEN")
    joinrole := os.Getenv("DJOINROLEID")
    dguild := os.Getenv("DGUILD")

    gendiv := os.Getenv("DGENDIV")
    doublesdiv := os.Getenv("DDOUBLESDIV")
    return &structs.Discord{
        prefix,
        token,
        joinrole,
        dguild,
        gendiv,
        doublesdiv,
    }
}
