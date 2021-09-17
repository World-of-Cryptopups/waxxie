package commands

import (
	"fmt"
	"time"

	"github.com/World-of-Cryptopups/eosrpc.go"
	"github.com/World-of-Cryptopups/waxxie/lib"
	"github.com/World-of-Cryptopups/waxxie/utils"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/enescakir/emoji"
)

func (b *Bot) Info(m *gateway.MessageCreateEvent, args bot.RawArguments) (interface{}, error) {
	b.Ctx.Typing(m.ChannelID)

	wallet := string(args)

	// discord id
	// discordId := m.Author.ID.String()
	acc, err := lib.CHAIN.GetAccount(eosrpc.AccountProps{AccountName: wallet})
	if err != nil {
		return utils.FailedMessage("Error trying to fetch account details!", err)
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

	embed := &discord.Embed{
		Author: &discord.EmbedAuthor{
			Name: m.Author.Tag(),
			Icon: m.Author.AvatarURL(),
		},
		Thumbnail: &discord.EmbedThumbnail{
			URL: m.Author.AvatarURL(),
		},
		Title: acc.AccountName,
		Fields: []discord.EmbedField{
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
				Value:  fmt.Sprintf("%s/%s", utils.ByteCountSI(cpuUsage), utils.ByteCountSI(cpuQuota)),
				Inline: true,
			},
			{
				Name:   fmt.Sprintf("%v NET", emoji.GlobeWithMeridians),
				Value:  fmt.Sprintf("%s/%s", utils.ByteCountSI(netUsage), utils.ByteCountSI(netQuota)),
				Inline: true,
			},
		},

		Footer: &discord.EmbedFooter{
			Text: "2021 | Waxxie - WoC",
		},
		Timestamp: discord.Timestamp(time.Now()),
	}

	return embed, nil
}
