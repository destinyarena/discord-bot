package roles

import (
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
    "github.com/arturoguerra/d2arena/internal/structs"
    "github.com/arturoguerra/d2arena/internal/config"
    "net/http"
    "strings"
    "gopkg.in/go-playground/validator.v9"
    "fmt"
)


var hubs []*structs.DiscordReaction

func init() {
    discord := config.LoadDiscord()
    hubs = discord.Reactions
}

func checkHub(hubid string, guid string) bool {
    return false
}


func sendInvites(s *discordgo.Session, guildid string, p *structs.RolesPayload, cfg *structs.Discord) {
    message := "Invite links to join FACEIT Hubs\n\n"
    send := false
    for _, hub := range hubs {
        if inhub := checkHub(hub.HubID, p.Faceit); inhub == false {
            fmt.Println(hub)
            if link, _ := sendLink(hub.HubID); link != "" {
                message += strings.Replace(hub.Format, "{invite}", link, 1) + "\n"
                send = true
            }
        }
    }

    if send {
        channel, _ := s.UserChannelCreate(p.Discord)
        s.ChannelMessageSend(channel.ID, message)
    }

}

func New(s *discordgo.Session) echo.HandlerFunc {
    discord := config.LoadDiscord()
    return func(c echo.Context) error {
        g, err := s.Guild(discord.GuildID)
        if err != nil {
            fmt.Println(err)
            return c.String(500, "Well shit we fucked up")
        }

        payload := new(structs.RolesPayload)
        if err := c.Bind(payload); err != nil {
            return c.String(http.StatusBadRequest, "Error invalid payload")
        }

        v := validator.New()

        if err = v.Struct(payload); err != nil {
            return c.String(http.StatusBadRequest, "Error invalid payload")
        }

        //sendInvites(s, g.ID, payload, discord)
        channel, _ := s.UserChannelCreate(payload.Discord)
        s.ChannelMessageSend(channel.ID, "Please go back to discord to finish registration")
        s.GuildMemberRoleAdd(g.ID, payload.Discord, discord.RegistrationRoleID)
        return c.String(http.StatusOK, "Roles have been assigned")
    }
}
