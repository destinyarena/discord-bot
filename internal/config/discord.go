package config

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "os"
)

func LoadDiscord() *structs.Discord {
    prefix := os.Getenv("PREFIX")
    token := os.Getenv("DTOKEN")
    joinrole := os.Getenv("DJOINROLEID")
    dguild := os.Getenv("DGUILD");

    div1 := os.Getenv("DDIV1");
    div2 := os.Getenv("DDIV2");
    div3 := os.Getenv("DDIV3");
    divD := os.Getenv("DDIVD");
    divS := os.Getenv("DDIVS");
    return &structs.Discord{
        prefix,
        token,
        joinrole,
        dguild,
        div1,
        div2,
        div3,
        divD,
        divS,
    }
}
