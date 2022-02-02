package commands

import (
	"TarkovBot/commands/embed"
	"TarkovBot/config"
	"log"

	"github.com/bwmarrin/discordgo"
)

func Reload(m *discordgo.MessageCreate, s *discordgo.Session, args []string) {
	// get discord ID

	for _, v := range config.BotSettings.AuthorizedIds {
		if v == m.Author.ID {
			// found authorized ID
			if err := config.LoadFile(); err != nil {
				log.Fatal(err.Error())
			}
			// Load Initial bot data
			if err := config.LoadSettings(); err != nil {
				log.Fatal(err.Error())
			}

			s.ChannelMessageSendEmbed(m.ChannelID, embed.NewGenericEmbed("Success", "reloaded `./assets/ammo.json` **and** `./assets/settings.json`"))
			return
		}
	}

	// didnt find authorized ID

	s.ChannelMessageSendEmbed(m.ChannelID, embed.NewErrorEmbed("Error", "Unauthorized!"))

}
