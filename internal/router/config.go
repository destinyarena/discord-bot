package router

type Config struct {
    Prefix string
    Token string
}

func NewConfig(prefix string, token string) *Config {
    return &Config{
        prefix,
        token,
    }
}
