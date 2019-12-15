package config

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "os"
)

func LoadFaceit() *structs.Faceit {
    utoken := os.Getenv("FACEIT_USER_TOKEN")
    apitoken := os.Getenv("FACEIT_API_TOKEN")

    gendiv := os.Getenv("FACEIT_GENERAL_DIV")
    doublesdiv := os.Getenv("FACEIT_DOUBLES_DIV")

    return &structs.Faceit{
        utoken,
        apitoken,
        gendiv,
        doublesdiv,
    }
}
