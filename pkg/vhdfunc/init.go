package vhdfunc

import (
	"flag"
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
