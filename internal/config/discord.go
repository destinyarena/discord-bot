package config

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "strings"
    "os"
)

func LoadDiscord() *structs.Discord {
    prefix := os.Getenv("DISCORD_PREFIX")
    token := os.Getenv("DISCORD_TOKEN")

    dguild := os.Getenv("DISCORD_GUILD_ID")
    logsid := os.Getenv("DISCORD_LOGS_CHANNEL")
    joinrole := os.Getenv("DISCORD_JOIN_ROLE_ID")
    bannedrole := os.Getenv("DISCORD_BANNED_ROLE_ID")

    owners := strings.Split(os.Getenv("DISCORD_OWNERS"), " ")
    staffroles := strings.Split(os.Getenv("DISCORD_STAFF_ROLES"), " ")

    registrationrole := os.Getenv("DISCORD_REGISTRATION_ROLE_ID")


    return &structs.Discord{
        Prefix: prefix,
        Token: token,

        GuildID: dguild,
        LogsID: logsid,
        JoinRoleID: joinrole,
        BannedRoleID: bannedrole,

        Owners: owners,
        StaffRoles: staffroles,

        RegistrationRoleID: registrationrole,
        Hubs: LoadHubs(),
        ReactionRoles: LoadReactionRoles(),
    }
}
