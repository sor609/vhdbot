package vhdfunc

import (
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// This function is called every time a new message is created on any channel
// that the authenticated bot has access to.
func ListenAndReply(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// help function
	if m.Content == CmdChar+"help" {
		helpMsg := PrintHelp()
		s.ChannelMessageSend(m.ChannelID, helpMsg)
	}

	// load up intents here
	msg, _ := regexp.MatchString("(?i)hey "+s.State.User.Username, m.Content)

	if msg {
		s.ChannelMessageSend(m.ChannelID, "What's up "+m.Author.Username+"?")
	}

	if m.Content == CmdChar+"start" || m.Content == CmdChar+"stop" || m.Content == CmdChar+"status" || m.Content == CmdChar+"rebot" {
		CtrlGame(strings.TrimPrefix(m.Content, CmdChar), s, m)
	}
	// load up functions here

}
