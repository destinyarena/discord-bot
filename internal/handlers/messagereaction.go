package handlers

import (
	"github.com/andersfylling/disgord"
)

type GenericReaction struct {
	ChannelID disgord.Snowflake
	MessageID disgord.Snowflake
	UserID    disgord.Snowflake
	Emoji     *disgord.Emoji
}

func (h *handler) reactionroles(s disgord.Session, reaction GenericReaction, addrole bool) {

	for _, rr := range h.Bot.Config.ReactionRoles {
		emoji := reaction.Emoji.ID.String()
		if len(emoji) == 0 {
			emoji = reaction.Emoji.Name
		}

		if rr.EmojiID == emoji && rr.MessageID == reaction.MessageID.String() {
			if addrole {
				s.Guild(disgord.ParseSnowflakeString(h.Bot.Config.GuildID)).Member(reaction.MessageID).AddRole(disgord.ParseSnowflakeString(rr.RoleID))
			} else {
				s.Guild(disgord.ParseSnowflakeString(h.Bot.Config.GuildID)).Member(reaction.MessageID).RemoveRole(disgord.ParseSnowflakeString(rr.RoleID))
			}
		}
	}
}
