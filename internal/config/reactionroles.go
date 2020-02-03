package config

import (
    "os"
    "github.com/arturoguerra/d2arena/internal/structs"
)

func LoadReactionRoles() []*structs.ReactionRole {
    pc := &structs.ReactionRole{
        MessageID: os.Getenv("RR_PC_MESSAGE_ID"),
        EmojiID:   os.Getenv("RR_PC_EMOJI_ID"),
        RoleID:    os.Getenv("RR_PC_ROLE_ID"),
    }

    ps4 := &structs.ReactionRole{
        MessageID: os.Getenv("RR_PS4_MESSAGE_ID"),
        EmojiID:   os.Getenv("RR_PS4_EMOJI_ID"),
        RoleID:    os.Getenv("RR_PS4_ROLE_ID"),
    }


    xbox := &structs.ReactionRole{
        MessageID: os.Getenv("RR_XBOX_MESSAGE_ID"),
        EmojiID:   os.Getenv("RR_XBOX_EMOJI_ID"),
        RoleID:    os.Getenv("RR_XBOX_ROLE_ID"),
    }

    return []*structs.ReactionRole{
        pc,
        ps4,
        xbox,
    }
}
