package commands

import (
	"TarkovBot/config"
	"log"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name        string
	Description string
	Alias       []string
	Execute     func(m *discordgo.MessageCreate, s *discordgo.Session, args []string)
}

var (
	Commands = make(map[string]*Command)
	Mutex    sync.Mutex
)

func Load() {
	register(&Command{
		Name:        "searchammo",
		Description: "search types of ammo",
		Alias:       []string{"sa"},
		Execute:     SA,
	})
	register(&Command{
		Name:        "reload",
		Description: "reload configuration files",
		Alias:       []string{"rl"},
		Execute:     Reload,
	})

	register(&Command{
		Name:        "help",
		Description: "gets this help menu",
		Alias:       []string{"info"},
		Execute:     Help,
	})
}

func register(cmd *Command) {
	// check for conflicting Command Names
	if _, found := Commands[cmd.Name]; found {
		log.Fatal("conflicting command names")
	}

	// check for conflicting command names through alias
	for _, cmds := range Commands {
		for _, v := range cmds.Alias {
			if cmd.Name == v {
				log.Fatal("conflicting command alias")
			}
		}
	}
	// if not found, register command in map
	// Lock a go concurrent
	Mutex.Lock()
	Commands[cmd.Name] = cmd
	Mutex.Unlock()

	log.Println("[CMD/ADD]: Name:", cmd.Name, " Description:", cmd.Description)

}

func Get(cmd string) (*Command, bool) {
	for _, v := range Commands {
		if cmd == config.BotSettings.Prefix+v.Name {
			return v, true
		}
	}
	// if not found search through command alias
	for _, v := range Commands {
		for _, alias := range v.Alias {
			if cmd == config.BotSettings.Prefix+alias {
				return v, true
			}
		}
	}
	return nil, false
}
