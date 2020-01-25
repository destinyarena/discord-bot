package config

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "os"
)

func LoadDiscord() *structs.Discord {
    prefix := os.Getenv("DISCORD_PREFIX")
    token := os.Getenv("DISCORD_TOKEN")

    dguild := os.Getenv("DISCORD_GUILD_ID")
    bannedrole := os.Getenv("DISCORD_BANNED_ROLE_ID")
    joinrole := os.Getenv("DISCORD_JOIN_ROLE_ID")
    staffrole := os.Getenv("DISCORD_STAFF_ROLE_ID")

    registrationrole := os.Getenv("DISCORD_REGISTRATION_ROLE_ID")


    rulesrole := os.Getenv("DISCORD_RULES_ROLE_ID")
    rulesmessage := os.Getenv("DISCORD_RULES_MESSAGE_ID")
    rulesemojiid := os.Getenv("DISCORD_RULES_EMOJI_ID")

    logsid := os.Getenv("DISCORD_LOGS_CHANNEL")

    hubs := LoadHubs()
    return &structs.Discord{
        Prefix: prefix,
        Token: token,
        GuildID: dguild,
        JoinRoleID: joinrole,
        StaffRoleID: staffrole,
        RegistrationRoleID: registrationrole,
        RulesRoleID: rulesrole,
        RulesMessageID: rulesmessage,
        RulesEmojiID: rulesemojiid,
        Hubs: hubs,
        BannedRoleID: bannedrole,
        LogsID: logsid,
    }
}
