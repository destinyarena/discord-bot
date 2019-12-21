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
        embed := &discordgo.MessageEmbed{
            Description: "Please click this [link](https://discordapp.com/channels/650109209610027034/657733307353792524/657760138282795018) to get your hub invites!",
        }
        if _, err := s.ChannelMessageSendEmbed(channel.ID, embed); err == nil {
            s.GuildMemberRoleAdd(g.ID, payload.Discord, discord.RegistrationRoleID)
        } else {
            embed := &discordgo.MessageEmbed{
                Title: "403: Forbidden",
                Description: fmt.Sprintf("Error sending hub channel to <@%s> please contact them", payload.Discord),
            }

            s.ChannelMessageSendEmbed(discord.LogsID, embed)
        }

        return c.String(http.StatusOK, "Roles have been assigned")
    }}
