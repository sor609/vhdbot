package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/sor609/vhdbot/pkg/vhdfunc"
)

var (
	BotToken string
	cmdChar  string
	adminMsg string
	adminChn string
	BotChans bool
)

func init() {

	flag.StringVar(&BotToken, "t", "", "Bot Token")
	flag.StringVar(&cmdChar, "c", "!", "Bot Command Character")
	flag.BoolVar(&BotChans, "l", false, "List Bot channels")
	flag.StringVar(&adminMsg, "m", "", "Send a Single Message <Data>")
	flag.StringVar(&adminChn, "ac", "", "Single Message Channel <Channel ID>")
	flag.Parse()

}

// this sets up bot and runs it
func startSession() {

	// Starting a new session
	session, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	// event handler for incoming messages
	session.AddHandler(listenAndReply)
	session.Identify.Intents = discordgo.IntentsGuildMessages

	err = session.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	// get list of bot channels
	if BotChans {
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
	if len(adminMsg) > 0 && len(adminChn) > 0 {
		session.ChannelMessageSend(adminChn, adminMsg)
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

// This function is called every time a new message is created on any channel
// that the authenticated bot has access to.
func listenAndReply(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// help function
	if m.Content == cmdChar+"help" {
		helpMsg := vhdfunc.PrintHelp()
		s.ChannelMessageSend(m.ChannelID, helpMsg)
	}

	// load up intents here
	msg, _ := regexp.MatchString("(?i)hey "+s.State.User.Username, m.Content)

	if msg {
		s.ChannelMessageSend(m.ChannelID, "What's up "+m.Author.Username+"?")
	}

	if m.Content == cmdChar+"start" || m.Content == cmdChar+"stop" || m.Content == cmdChar+"status" || m.Content == cmdChar+"rebot" {
		vhdfunc.CtrlGame(strings.TrimPrefix(m.Content, cmdChar), s, m)
	}
	// load up functions here

}

func main() {
	startSession()
}
