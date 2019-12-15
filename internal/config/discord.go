package config

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "os"
)

func LoadDiscord() *structs.Discord {
    prefix := os.Getenv("PREFIX")
    token := os.Getenv("DTOKEN")
    joinrole := os.Getenv("DJOINROLEID")
    dividerrole := os.Getenv("DISCORD_DIVIDER_ROLE")
    dguild := os.Getenv("DGUILD")

    gendiv := os.Getenv("DGENERALDIV")
    doublesdiv := os.Getenv("DDOUBLESDIV")
    return &structs.Discord{
        prefix,
        token,
        joinrole,
        dividerrole,
        dguild,
        gendiv,
        doublesdiv,
    }
}
