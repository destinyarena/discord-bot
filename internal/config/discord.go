package config

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "os"
)

func LoadDiscord() *structs.Discord {
    prefix := os.Getenv("PREFIX")
    token := os.Getenv("DTOKEN")
    div1 := os.Getenv("DIV1RID")
    div2 := os.Getenv("DIV2RID")
    div3 := os.Getenv("DIV3RID")

    return &structs.Discord{
        prefix,
        token,
        div1,
        div2,
        div3,
    }
}
