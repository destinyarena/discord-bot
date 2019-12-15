package moderation

import (
    "github.com/arturoguerra/d2arena/internal/router"
)

func getProfile(ctx *router.Context) {
    if len(ctx.Mentions) != 1 {
        return
    }

    if ctx.ChannelID != StaffChannelID {
        return
    }

    uid := ctx.Mentions[0].ID

    profile, err := fetchProfile(uid)
    if err != nil {
        ctx.Session.ChannelMessageSend(ctx.ChannelID, "Error fetching user profile")
        return
    }

    ctx.Session.ChannelMessageSend(ctx.ChannelID, generateProfile(ctx, profile))
}

func generateProfile (ctx *router.Context, profile *Profile) string {
    return "Faceit Name: " + profile.FaceitName + "\n" + "FaceitID: " + profile.FaceitGuid + "\n" + "Discord ID:" + profile.DiscordID + "\n"
}
