package config

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "os"
)

func LoadDiscord() *structs.Discord {
    prefix := os.Getenv("DISCORD_PREFIX")
    token := os.Getenv("DISCORD_TOKEN")
    joinrole := os.Getenv("DISCORD_JOIN_ROLE_ID")
    dguild := os.Getenv("DISCORD_GUILD_ID")

    faceit := os.Getenv("DISCORD_FACEIT_DIV_ROLE")
    return &structs.Discord{
        prefix,
        token,
        joinrole,
        dguild,
        faceit,
    }
}
