package vhdfunc

import (
	"flag"
	"fmt"
	"os"
)

var (
	BotToken string
	CmdChar  string
	AdminMsg string
	AdminChn string
	JFile    string
	BotChans bool
)

func VhdInit() {

	// read command line switches
	// we will also check for JSON config
	flag.StringVar(&BotToken, "t", "", "Bot Token")
	flag.StringVar(&CmdChar, "c", "!", "Bot Command Character")
	flag.StringVar(&JFile, "f", "", "JSON config")
	flag.BoolVar(&BotChans, "l", false, "List Bot channels")
	flag.StringVar(&AdminMsg, "m", "", "Send a Single Message <Data>")
	flag.StringVar(&AdminChn, "ac", "", "Single Message Channel <Channel ID>")
	flag.Parse()

	if JFile == "" && BotToken == "" {
		fmt.Println("You have not provided a token")
		os.Exit(1)
	}

}
