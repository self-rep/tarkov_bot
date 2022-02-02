package commands

import (
	"TarkovBot/commands/embed"
	"TarkovBot/config"
	"fmt"
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/bwmarrin/discordgo"
)

func Help(m *discordgo.MessageCreate, s *discordgo.Session, args []string) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{
				Align: simpletable.AlignCenter,
				Text:  "#",
			},
			{
				Align: simpletable.AlignCenter,
				Text:  "Name",
			},
			{
				Align: simpletable.AlignCenter,
				Text:  "Description",
			},
			{
				Align: simpletable.AlignCenter,
				Text:  "Alias",
			},
		},
	}

	var i int
	for _, v := range Commands {
		r := []*simpletable.Cell{
			{
				Align: simpletable.AlignRight,
				Text:  fmt.Sprint(i + 1),
			},
			{
				Text: config.BotSettings.Prefix + v.Name,
			},
			{
				Text: v.Description,
			},
			{
				Text: strings.Join(v.Alias, ", "),
			},
		}
		table.Body.Cells = append(table.Body.Cells, r)
		i++
	}

	table.SetStyle(simpletable.StyleCompactClassic)
	s.ChannelMessageSendEmbed(m.ChannelID, embed.NewGenericEmbed("[+] Commands List [+]", fmt.Sprintf("```%s```", table.String())))
}
