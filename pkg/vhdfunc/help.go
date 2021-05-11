package vhdfunc

// Self-explanatory, print help if user asks bot
// for command list
func PrintHelp() string {
	help := "```ini\n"
	help += "[Bot Commands]\n"
	help += `
	start  - start Valheim server
	stop   - stop Valheim server
	status - get Valheim server status
	rebot  - reboot Bot`
	help += "```"

	return help
}
