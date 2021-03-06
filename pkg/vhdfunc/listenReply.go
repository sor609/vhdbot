package vhdfunc

import (
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

	// game controls
	if m.Content == CmdChar+"start" || m.Content == CmdChar+"stop" || m.Content == CmdChar+"status" || m.Content == CmdChar+"rebot" {
		CtrlGame(strings.TrimPrefix(m.Content, CmdChar), s, m)
	}
}
