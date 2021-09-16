package main

import (
	"log"
	"os"

	"github.com/World-of-Cryptopups/waxxie/commands"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var token = os.Getenv("TOKEN")

	if token == "" {
		log.Fatal("Missing TOKEN!")
	}

	commands := &commands.Bot{}
	bot.Run(token, commands, func(c *bot.Context) error {
		c.HasPrefix = bot.NewPrefix("?")
		c.EditableCommands = true

		// do not show unknown commands error
		c.SilentUnknown.Command = true
		c.SilentUnknown.Subcommand = true

		c.AddIntents(gateway.IntentDirectMessages) // for dm messages
		c.AddIntents(gateway.IntentGuildMessages)  // for guild or server messages
		c.AddIntents(gateway.IntentGuilds)

		c.Gateway.Identifier.IdentifyData = gateway.IdentifyData{
			Token: c.Token,
			Presence: &gateway.UpdateStatusData{
				Activities: []discord.Activity{
					{
						Name: "Wax Chain API",
						Type: discord.WatchingActivity,
					},
				},
			},
		}

		return nil
	})
}
