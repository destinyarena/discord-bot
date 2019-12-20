package config

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "os"
)

func LoadFaceit() *structs.Faceit {
    utoken := os.Getenv("FACEIT_USER_TOKEN")
    apitoken := os.Getenv("FACEIT_API_TOKEN")

    gendiv := os.Getenv("DIVISION_ARENA_ID")
    doublesdiv := os.Getenv("DIVISION_DOUBLES_ID")

    return &structs.Faceit{
        utoken,
        apitoken,
        gendiv,
        doublesdiv,
    }
}
