package moderation

import (
    "github.com/arturoguerra/d2arena/internal/router"
    "strings"
)

func getProfile(ctx *router.Context) {
    if ctx.ChannelID != StaffChannelID {
        return
    }

    var uid string

    if len(ctx.Mentions) > 0 {
        uid = ctx.Mentions[0].ID
    } else {
        split := strings.Split(ctx.Content, " ")
        uid = strings.Join(split[2:], " ")
    }

    if uid == "" {
        return
    }

    idtype := sortProfileId(uid)

    profile, err := fetchProfile(uid, idtype)
    if err != nil {
        ctx.Session.ChannelMessageSend(ctx.ChannelID, "Error fetching user profile")
        return
    }

    ctx.Session.ChannelMessageSend(ctx.ChannelID, generateProfile(ctx, profile))
}

func generateProfile (ctx *router.Context, profile *Profile) string {
    return "Faceit Name: " + profile.FaceitName + "\n" + "FaceitID: " + profile.FaceitGuid + "\n" + "Discord ID: " + profile.DiscordID + "\n"
}
