package handlers

import "github.com/andersfylling/disgord"

func (h *handler) OnMessageReactionAdd() func(disgord.Session, *disgord.MessageReactionAdd) {
	return func(s disgord.Session, mra *disgord.MessageReactionAdd) {
		r := GenericReaction{
			ChannelID: mra.ChannelID,
			MessageID: mra.MessageID,
			UserID:    mra.UserID,
			Emoji:     mra.PartialEmoji,
		}
		h.reactionroles(s, r, true)
	}
}
