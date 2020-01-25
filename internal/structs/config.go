package structs

// Configuration

type (
    Config struct {
        Faceit *Faceit
        Discord *Discord
    }

    Faceit struct {
        UserToken string
        ApiToken string
    }

    Discord struct {
        // Bot Stuff
        Prefix string
        Token string

        // Guild Stuff
        GuildID string
        LogsID string
        JoinRoleID string
        StaffRoleID string
        BannedRoleID string

        // Reaction roles stuff
        RegistrationRoleID string
        RulesRoleID string
        RulesMessageID string
        RulesEmojiID string
        Hubs []*Hub
    }

    Hub struct {
        Format string
        HubID string
        RoleID string
        EmojiID string
        MessageID string
        Main bool
        SkillLvl int
    }
)
