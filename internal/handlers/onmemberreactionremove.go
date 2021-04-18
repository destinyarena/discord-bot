package handlers

import "github.com/andersfylling/disgord"

func (h *handler) OnMessageReactionRemove() func(disgord.Session, *disgord.MessageReactionRemove) {
	return func(s disgord.Session, mrr *disgord.MessageReactionRemove) {
		r := GenericReaction{
			ChannelID: mrr.ChannelID,
			MessageID: mrr.MessageID,
			UserID:    mrr.UserID,
			Emoji:     mrr.PartialEmoji,
		}
		h.reactionroles(s, r, false)
	}
}
