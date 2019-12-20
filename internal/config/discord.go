package config

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "os"
)

func loadReactions() []*structs.DiscordReaction {
    arenaformat := os.Getenv("DIVISION_ARENA_FORMAT")
    arenaid := os.Getenv("DIVISION_ARENA_ID")
    arenaroleid := os.Getenv("DISCORD_DIVISION_ARENA_ROLE_ID")
    arenaemojiid := os.Getenv("DISCORD_DIVIONS_ARENA_EMOJI_ID")

    doublesformat := os.Getenv("DIVISION_DOUBLES_FORMAT")
    doublesid := os.Getenv("DIVISION_DOUBLES_ID")
    doublesroleid := os.Getenv("DISCORD_DIVISION_DOUBLES_ROLE_ID")
    doublesemojiid := os.Getenv("DISCORD_DIVISION_DOUBLES_EMOJI_ID")

    return []*structs.DiscordReaction{
        &structs.DiscordReaction{
            arenaformat,
            arenaid,
            arenaroleid,
            arenaemojiid,
            1,
        },
        &structs.DiscordReaction{
            doublesformat,
            doublesid,
            doublesroleid,
            doublesemojiid,
            1,
        },
    }
}

func LoadDiscord() *structs.Discord {
    prefix := os.Getenv("DISCORD_PREFIX")
    token := os.Getenv("DISCORD_TOKEN")

    joinrole := os.Getenv("DISCORD_JOIN_ROLE_ID")
    dguild := os.Getenv("DISCORD_GUILD_ID")
    staffrole := os.Getenv("DISCORD_STAFF_ROLE_ID")

    registrationrole := os.Getenv("DISCORD_REGISTRATION_ROLE_ID")
    registrationmsg := os.Getenv("DISCORD_REGISTRATION_MESSAGE_ID")

    invitesmsgid := os.Getenv("DISCORD_INVITES_MESSAGE_ID")
    reactions := loadReactions()
    return &structs.Discord{
        Prefix: prefix,
        Token: token,
        GuildID: dguild,
        JoinRoleID: joinrole,
        StaffRoleID: staffrole,
        RegistrationMessageID: registrationmsg,
        RegistrationRoleID: registrationrole,
        InvitesMessageID: invitesmsgid,
        Reactions: reactions,
    }
}
