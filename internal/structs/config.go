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
        Roles *Roles
    }

    Roles struct {
        Div1 int
        Div2 int
        Div3 int
    }
)
