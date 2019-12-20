package handlers

import (
  //  "github.com/arturoguerra/d2arena/internal/structs"
//    "github.com/arturoguerra/d2arena/internal/config"
    "github.com/bwmarrin/discordgo"
    "fmt"
)
func OnMessageReactionAdd(s *discordgo.Session, mr *discordgo.MessageReactionAdd) {
    fmt.Println(mr)
    fmt.Println(mr.Emoji)
}
