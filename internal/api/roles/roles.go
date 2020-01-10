package roles

import (
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
    "github.com/arturoguerra/d2arena/internal/structs"
    "github.com/arturoguerra/d2arena/internal/config"
    "net/http"
    "errors"
    "strings"
    "gopkg.in/go-playground/validator.v9"
    "fmt"
)


var hubs []*structs.Hub

func init() {
    discord := config.LoadDiscord()
    hubs = discord.Hubs
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

        u, err := s.User(payload.Discord)
        if err != nil {
            return errors.New("User not found")
        }

        channel, err := s.UserChannelCreate(payload.Discord)
        if err != nil {
            return errors.New("Channel not found")
        }

        embed := &discordgo.MessageEmbed{
            Description: "Please click this [link](https://discordapp.com/channels/650109209610027034/657733307353792524/665277347909599232) to get your hub invites!",
        }

        if _, err := s.ChannelMessageSendEmbed(channel.ID, embed); err == nil {
            s.GuildMemberRoleAdd(g.ID, payload.Discord, discord.RegistrationRoleID)
            embed := &discordgo.MessageEmbed{
                Title: "Hub invite link",
                Description: fmt.Sprintf("Sent hubs channel link to <@%s>(`%s#%s`)", payload.Discord, u.Username, u.Discriminator),
            }

            s.ChannelMessageSendEmbed(discord.LogsID, embed)
        } else {
            embed := &discordgo.MessageEmbed{
                Title: "403: Forbidden",
                Description: fmt.Sprintf("Error sending hub channel to <@%s>(`%s#%s`) please contact them", payload.Discord, u.Username, u.Discriminator),
            }

            s.ChannelMessageSendEmbed(discord.LogsID, embed)
        }

        return c.String(http.StatusOK, "Roles have been assigned")
    }
}
