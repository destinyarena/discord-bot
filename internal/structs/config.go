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
        Prefix string
        Token string
        JoinRoleID string
        DividerRoleID string
        GuildID string
        GeneralDiv string
        DoublesDiv string
    }
)
