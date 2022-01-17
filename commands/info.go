package commands

import (
	"fmt"
	"time"

	"github.com/World-of-Cryptopups/eosrpc.go"
	"github.com/World-of-Cryptopups/minidis"
	"github.com/World-of-Cryptopups/waxxie/lib"
	"github.com/World-of-Cryptopups/waxxie/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/enescakir/emoji"
)

var InfoCMD = &minidis.SlashCommandProps{
	Command:     "info",
	Description: "Get wallet info of an eosio account.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "wallet",
			Description: "Wallet to fetch details.",
			Required:    true,
		},
	},
	Execute: func(c *minidis.SlashContext) error {
		c.DeferReply(false)

		wallet := c.Options["wallet"].StringValue()

		// discordId := m.Author.ID.String()
		acc, err := lib.CHAIN.GetAccount(eosrpc.AccountProps{AccountName: wallet})
		if err != nil {
			return c.Edit("Error trying to fetch account details!")
		}

		// ram usage
		ramUsage, _ := acc.RAMUsage.Int64()
		ramQuota, _ := acc.RAMQuota.Int64()

		// cpu usage
		cpuUsage, _ := acc.CPULimit.Used.Int64()
		cpuQuota, _ := acc.CPULimit.Max.Int64()

		// net usage
		netUsage, _ := acc.NetLimit.Used.Int64()
		netQuota, _ := acc.NetLimit.Max.Int64()

		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    c.Author.String(),
				IconURL: c.Author.AvatarURL(""),
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: c.Author.AvatarURL(""),
			},
			Title: acc.AccountName,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Available Balance",
					Value: fmt.Sprintf("%v **%s**", emoji.Coin, acc.CoreLiquidBalance),
				},
				{
					Name:  "\u200b",
					Value: "\u200b",
				},
				{
					Name:   fmt.Sprintf("%v RAM", emoji.ComputerDisk),
					Value:  fmt.Sprintf("%s/%s", utils.ByteCountSI(ramUsage), utils.ByteCountSI(ramQuota)),
					Inline: true,
				},
				{
					Name:   fmt.Sprintf("%v CPU", emoji.DesktopComputer),
					Value:  fmt.Sprintf("%s/%s", utils.ByteCountSI(cpuUsage), utils.ByteCountSI(cpuQuota)), // this is wrong, should be ms not filesize
					Inline: true,
				},
				{
					Name:   fmt.Sprintf("%v NET", emoji.GlobeWithMeridians),
					Value:  fmt.Sprintf("%s/%s", utils.ByteCountSI(netUsage), utils.ByteCountSI(netQuota)),
					Inline: true,
				},
			},

			Footer: &discordgo.MessageEmbedFooter{
				Text: "2022 | Waxxie - WoC",
			},
			Timestamp: time.Now().Format(time.RFC3339),
		}

		return c.EditC(minidis.EditProps{
			Content: "",
			Embeds: []*discordgo.MessageEmbed{
				embed,
			},
		})
	},
}
