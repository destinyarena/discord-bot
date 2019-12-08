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
        Prefix string
        Token string
        JoinRoleID string
        GuildID string
        Div1 string
        Div2 string
        Div3 string
        DivD string
        DivS string
    }
)
