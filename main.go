package main

import (
	"log"
	"os"
	"strings"

	"github.com/World-of-Cryptopups/minidis"
	"github.com/World-of-Cryptopups/waxxie/commands"
	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	guilds := strings.Split(os.Getenv("GUILDS"), ",")

	// new minidis handler
	bot := minidis.New(os.Getenv("TOKEN"))

	bot.OnReady(func(s *discordgo.Session, i *discordgo.Ready) {
		log.Println("Bot is ready!")
	})

	// sync commands to these guilds only
	bot.SyncToGuilds(guilds...)

	bot.AddCommand(commands.InfoCMD)

	if err := bot.Run(); err != nil {
		log.Fatalln(err)
	}
}
