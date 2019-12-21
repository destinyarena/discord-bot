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
        GeneralDiv string
        DoublesDiv string
    }

    Discord struct {
        // Bot Stuff
        Prefix string
        Token string

        // Guild Stuff
        GuildID string
        JoinRoleID string
        StaffRoleID string
        BannedRoleID string

        // Reaction roles stuff
        RegistrationRoleID string
        RulesRoleID string
        RulesMessageID string
        RulesEmojiID string
        InvitesMessageID string
        InvitesAutoEmojiID string
        Reactions []*DiscordReaction
    }

    DiscordReaction struct {
        Format string
        HubID string
        RoleID string
        EmojiID string
        Main bool
        SkillLvl int
    }
)
