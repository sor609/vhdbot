package vhdfunc

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/valyala/fastjson"
)

var (
	JFile       string
	BotToken    string
	CmdChar     string
	AdminMsg    string
	AdminChn    string
	GameScript  string
	GameChannel string
	BotChans    bool
	VhdInts     = make(map[string]string)
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

	readJconfig()
}

func readJconfig() {
	if JFile != "" {
		j, err := os.Open(JFile)
		if err != nil {
			fmt.Println(err)
		}

		defer j.Close()

		data, _ := ioutil.ReadAll(j)

		var p fastjson.Parser
		cfg, err := p.ParseBytes(data)
		if err != nil {
			fmt.Println(err)
		}

		// command line token will override JSON
		if BotToken == "" {
			BotToken = string(cfg.GetStringBytes("config", "BotToken"))
		}

		GameScript = string(cfg.GetStringBytes("config", "GameScript"))
		GameChannel = string(cfg.GetStringBytes("config", "GameChannel"))

		// add all intents into a map so it's easier to parse

		cfg.GetObject("response").Visit(func(key []byte, val *fastjson.Value) {
			VhdInts[string(key)] = fmt.Sprintf("%s", val)
		})
	}
}
