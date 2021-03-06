package vhdfunc

import (
	"fmt"
	"os/exec"

	"github.com/bwmarrin/discordgo"
)

// function to control the game on linux systems
// bot has to run on the same server as where the game is
// and have permissions to execute the start/stop script
func CtrlGame(f string, s *discordgo.Session, m *discordgo.MessageCreate) {

	if GameScript == "" {
		fmt.Println("Game script not defined!")
		return
	} else {
		cmd := exec.Command(GameScript, f)
		if err := cmd.Run(); err != nil {
			fmt.Println("Game controller error: ", err)
		}
		s.ChannelMessageSend(m.ChannelID, "Executed: "+f)
	}
}
