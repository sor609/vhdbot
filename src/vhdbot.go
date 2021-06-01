package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/sor609/vhdbot/pkg/vhdfunc"
)

// do a quick init
func init() {
	vhdfunc.VhdInit()
}

// this sets up bot and runs it
func startSession() {

	// Starting a new session
	session, err := discordgo.New("Bot " + vhdfunc.BotToken)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	// event handler for incoming messages
	session.AddHandler(vhdfunc.ListenAndReply)
	session.Identify.Intents = discordgo.IntentsGuildMessages

	err = session.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	// get list of bot channels
	if vhdfunc.BotChans {
		for _, guild := range session.State.Guilds {
			channels, _ := session.GuildChannels(guild.ID)
			for _, c := range channels {
				if c.Type != discordgo.ChannelTypeGuildText {
					continue
				}
				fmt.Println(c.ID, c.Name)
			}
		}
		os.Exit(0)
	}

	// send single/admin message from cmdline
	// useful to notify game is going down or is down
	if len(vhdfunc.AdminMsg) > 0 && len(vhdfunc.AdminChn) > 0 {
		session.ChannelMessageSend(vhdfunc.AdminChn, vhdfunc.AdminMsg)
		os.Exit(0)
	}

	// Interrupt handler
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// cleanly close down the Discord session.
	session.Close()

}

func main() {
	startSession()
}
