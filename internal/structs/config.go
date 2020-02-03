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
        GuildID      string
        LogsID       string
        JoinRoleID   string
        BannedRoleID string

        // Permissions
        Owners       []string
        StaffRoles   []string

        // Reaction roles stuff
        RegistrationRoleID string
        Hubs []*Hub
        ReactionRoles []*ReactionRole
    }

    ReactionRole struct {
        EmojiID   string
        MessageID string
        RoleID    string
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
