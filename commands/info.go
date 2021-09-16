package commands

import "github.com/diamondburned/arikawa/v2/gateway"

func (b *Bot) Info(m *gateway.MessageCreateEvent) (interface{}, error) {
	b.Ctx.Typing(m.ChannelID)

	// discord id
	discordId := m.Author.ID.String()

	return discordId, nil
}
