package background

import (
    "github.com/bwmarrin/discordgo"
    "time"
    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/arturoguerra/d2arena/internal/checkmark"
)


func checkroles(roles []string, roleid string) bool {
    for _, role := range roles {
        if role == roleid {
            return false
        }
    }

    return true
}

func filter(members []*discordgo.Member, rid string) []*discordgo.Member {
    var mlist []*discordgo.Member
    for _, member := range members {
        createdAt, _ := discordgo.SnowflakeTimestamp(member.User.ID)
        loc, _ := time.LoadLocation("UTC")
        now := time.Now().In(loc)
        diff := now.Sub(createdAt)

        days := int(diff.Hours()) / 24

        if  (days >= 7 && checkroles(member.Roles, rid)) {
            mlist = append(mlist, member)
        }
    }

    return mlist
}

func UpdateRoles(s *discordgo.Session) {
    discord := config.LoadDiscord()
    rid := discord.JoinRoleID
    for true {
        guild, err := s.Guild("650109209610027034")
        if err == nil {
            members := filter(guild.Members, rid)
            for _, member := range members {
                checkmark.Member(s, member)
            }
        }

        time.Sleep(120 * time.Second)
    }
}
