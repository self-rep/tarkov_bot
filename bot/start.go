package bot

import (
	"TarkovBot/commands"
	"TarkovBot/config"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func NewBot() {
	disc, err := discordgo.New("Bot " + config.BotSettings.Token)
	if err != nil {
		log.Fatal(err.Error())
	}
	disc.AddHandler(MessageHandler)

	err = disc.Open()
	if err != nil {
		log.Fatal("Failed to open connection to bot!")
	}
	log.Println("logged into:", disc.State.User.Username+"#"+disc.State.User.Discriminator)
	waitforsig := make(chan os.Signal, 1)
	signal.Notify(waitforsig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-waitforsig

}

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// TODO: Create suitable message handler
	if m.Author.ID == s.State.User.ID {
		return
	}

	cmdlist := strings.Split(m.Content, " ") // split message into array

	if len(cmdlist) == 0 {
		// if somehow we have a empty message skip over it
		return
	}

	if strings.HasPrefix(cmdlist[0], config.BotSettings.Prefix) {
		command, ok := commands.Get(cmdlist[0])
		if !ok {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("```Command: %s isnt recognized```", cmdlist[0]))
			return
		}
		commands.Commands[command.Name].Execute(m, s, cmdlist)
	}
}
