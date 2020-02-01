package roles

import (
    "github.com/bwmarrin/discordgo"
    "github.com/labstack/echo"
    "github.com/arturoguerra/d2arena/internal/config"
    "net/http"
    "errors"
    "fmt"
)



func New(s *discordgo.Session) echo.HandlerFunc {
    discord := config.LoadDiscord()
    return func(c echo.Context) error {
        g, err := s.Guild(discord.GuildID)
        if err != nil {
            fmt.Println(err)
            return c.String(500, "Well shit we fucked up")
        }

        uid := c.Param("id")


        u, err := s.User(uid)
        if err != nil {
            return errors.New("User not found")
        }

        channel, err := s.UserChannelCreate(uid)
        if err != nil {
            return errors.New("Channel not found")
        }

        embed := &discordgo.MessageEmbed{
            Description: "Please click this [link](https://discordapp.com/channels/650109209610027034/657733307353792524/665277347909599232) to get your hub invites!",
        }

        if _, err := s.ChannelMessageSendEmbed(channel.ID, embed); err == nil {
            s.GuildMemberRoleAdd(g.ID, uid, discord.RegistrationRoleID)
            embed := &discordgo.MessageEmbed{
                Title: "Hub invite link",
                Description: fmt.Sprintf("Sent hubs channel link to <@%s>(`%s#%s`)", uid, u.Username, u.Discriminator),
            }

            s.ChannelMessageSendEmbed(discord.LogsID, embed)
        } else {
            embed := &discordgo.MessageEmbed{
                Title: "403: Forbidden",
                Description: fmt.Sprintf("Error sending hub channel to <@%s>(`%s#%s`) please contact them", uid, u.Username, u.Discriminator),
            }

            s.ChannelMessageSendEmbed(discord.LogsID, embed)
        }

        return c.String(http.StatusOK, "Roles have been assigned")
    }
}
