package config

import (
    "github.com/arturoguerra/d2arena/internal/structs"
    "os"
)

func loadReactions() []*structs.DiscordReaction {
    arenaformat := os.Getenv("DIVISION_ARENA_FORMAT")
    arenaid := os.Getenv("DIVISION_ARENA_ID")
    arenaroleid := os.Getenv("DISCORD_DIVISION_ARENA_ROLE_ID")
    arenaemojiid := os.Getenv("DISCORD_DIVISION_ARENA_EMOJI_ID")

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
            true,
            0,
        },
        &structs.DiscordReaction{
            doublesformat,
            doublesid,
            doublesroleid,
            doublesemojiid,
            false,
            0,
        },
    }
}

func LoadDiscord() *structs.Discord {
    prefix := os.Getenv("DISCORD_PREFIX")
    token := os.Getenv("DISCORD_TOKEN")

    dguild := os.Getenv("DISCORD_GUILD_ID")
    bannedrole := os.Getenv("DISCORD_BANNED_ROLE_ID")
    joinrole := os.Getenv("DISCORD_JOIN_ROLE_ID")
    staffrole := os.Getenv("DISCORD_STAFF_ROLE_ID")

    registrationrole := os.Getenv("DISCORD_REGISTRATION_ROLE_ID")

    invitesmsgid := os.Getenv("DISCORD_INVITES_MESSAGE_ID")
    autoemoji := os.Getenv("DISCORD_INVITES_AUTO_EMOJI_ID")

    rulesrole := os.Getenv("DISCORD_RULES_ROLE_ID")
    rulesmessage := os.Getenv("DISCORD_RULES_MESSAGE_ID")
    rulesemojiid := os.Getenv("DISCORD_RULES_EMOJI_ID")

    logsid := os.Getenv("DISCORD_LOGS_CHANNEL")

    reactions := loadReactions()
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
        InvitesMessageID: invitesmsgid,
        InvitesAutoEmojiID: autoemoji,
        Reactions: reactions,
        BannedRoleID: bannedrole,
        LogsID: logsid,
    }
}
