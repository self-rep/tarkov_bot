package commands

import (
	"TarkovBot/commands/embed"
	"TarkovBot/search_engine"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// SA = Search Ammo

func SA(m *discordgo.MessageCreate, s *discordgo.Session, args []string) {
	// lower case search
	// $sa = args[0]

	joined := strings.ToLower(strings.Join(args[1:], " ")) // lower case the string to avoid sensitive search

	ammo, results := search_engine.Search(joined)
	if ammo != nil && results == nil {
		joinedguns := strings.Join(ammo.Weapons, ",")
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "Search Results",
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: ammo.ImageUrl,
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Penetration",
					Value:  ammo.Stats.Penetration,
					Inline: true,
				},
				{
					Name:   "Flesh Damage",
					Value:  ammo.Stats.FleshDamage,
					Inline: true,
				},
				{
					Name:   "Armour Damage",
					Value:  ammo.Stats.ArmourDamage + "%",
					Inline: true,
				},
				{
					Name:   "Projectile Speed",
					Value:  ammo.Stats.ProjectileSpeed + "M/S",
					Inline: true,
				},
				{
					Name:   "Frag Chance",
					Value:  ammo.Stats.FragmentChance,
					Inline: true,
				},
				{
					Name:   "Richochet Chance",
					Value:  ammo.Stats.RichochetChance,
					Inline: true,
				},
			},
			Description: fmt.Sprintf("Compatible Weapons: **%s**", joinedguns),
		})
		return
	}
	if ammo == nil && results == nil {
		s.ChannelMessageSendEmbed(m.ChannelID, embed.NewErrorEmbed("Error", "Ammo Could Not Be Found..."))
		return
	}

	if results != nil && ammo == nil {
		var ammoarray []string
		for _, v := range results {
			ammoarray = append(ammoarray, "‣ "+v.Name)
		}

		sjoined := strings.Join(ammoarray, "\n")
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       "Items Found",
			Description: sjoined,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Search is not case sensitive",
			},
		})
	}
}
