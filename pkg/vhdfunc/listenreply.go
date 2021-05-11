package vhdfunc

import (
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// This function is called every time a new message is created on any channel
// that the authenticated bot has access to.
func listenAndReply(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// help function
	if m.Content == cmdChar+"help" {
		helpMsg := printHelp()
		s.ChannelMessageSend(m.ChannelID, helpMsg)
	}

	// load up intents here
	msg, _ := regexp.MatchString("(?i)hey "+s.State.User.Username, m.Content)

	if msg {
		s.ChannelMessageSend(m.ChannelID, "What's up "+m.Author.Username+"?")
	}

	if m.Content == cmdChar+"start" || m.Content == cmdChar+"stop" || m.Content == cmdChar+"status" || m.Content == cmdChar+"rebot" {
		ctrlGame(strings.TrimPrefix(m.Content, cmdChar), s, m)
	}
	// load up functions here

}
