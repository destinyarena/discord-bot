package config

import "github.com/arturoguerra/d2arena/internal/structs"

func LoadConfig() *structs.Config {
    discord := LoadDiscord()
    faceit := LoadFaceit()

    return &structs.Config{
        faceit,
        discord,
    }
}
