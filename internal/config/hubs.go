package config

import (
    "os"
    "github.com/arturoguerra/d2arena/internal/structs"
)

func LoadHubs() []*structs.Hub {
    pc := os.Getenv("DISCORD_INVITES_MESSAGE_ID")
    div1 := &structs.Hub{
        os.Getenv("HUB_1_FORMAT"),
        os.Getenv("HUB_1_ID"),
        os.Getenv("HUB_1_ROLE_ID"),
        os.Getenv("HUB_1_EMOJI_ID"),
        pc,
        true,
        8,
    }

    div2 := &structs.Hub{
        os.Getenv("HUB_2_FORMAT"),
        os.Getenv("HUB_2_ID"),
        os.Getenv("HUB_2_ROLE_ID"),
        os.Getenv("HUB_2_EMOJI_ID"),
        pc,
        true,
        4,
    }

    div3 := &structs.Hub{
        os.Getenv("HUB_3_FORMAT"),
        os.Getenv("HUB_3_ID"),
        os.Getenv("HUB_3_ROLE_ID"),
        os.Getenv("HUB_3_EMOJI_ID"),
        pc,
        true,
        0,
    }

    duel := &structs.Hub{
        os.Getenv("HUB_DUEL_FORMAT"),
        os.Getenv("HUB_DUEL_ID"),
        os.Getenv("HUB_DUEL_ROLE_ID"),
        os.Getenv("HUB_DUEL_EMOJI_ID"),
        pc,
        false,
        0,
    }

    double := &structs.Hub{
        os.Getenv("HUB_DOUBLES_FORMAT"),
        os.Getenv("HUB_DOUBLES_ID"),
        os.Getenv("HUB_DOUBLES_ROLE_ID"),
        os.Getenv("HUB_DOUBLES_EMOJI_ID"),
        pc,
        false,
        0,
    }

    stadium := &structs.Hub{
        os.Getenv("HUB_STADIUM_FORMAT"),
        os.Getenv("HUB_STADIUM_ID"),
        os.Getenv("HUB_STADIUM_ROLE_ID"),
        os.Getenv("HUB_STADIUM_EMOJI_ID"),
        pc,
        false,
        0,
    }

    pssolo := &structs.Hub{
        os.Getenv("HUB_PS_SOLO_FORMAT"),
        os.Getenv("HUB_PS_SOLO_ID"),
        os.Getenv("HUB_PS_SOLO_ROLE_ID"),
        os.Getenv("HUB_PS_SOLO_EMOJI_ID"),
        pc,
        true,
        0,
    }

    psteam := &structs.Hub{
        os.Getenv("HUB_PS_TEAM_FORMAT"),
        os.Getenv("HUB_PS_TEAM_ID"),
        os.Getenv("HUB_PS_TEAM_ROLE_ID"),
        os.Getenv("HUB_PS_TEAM_EMOJI_ID"),
        pc,
        true,
        0,
    }


    hubs := []*structs.Hub{
        div1,
        div2,
        div3,
        duel,
        double,
        stadium,
        pssolo,
        psteam,
    }

    return hubs
}
