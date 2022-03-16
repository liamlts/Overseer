package main

import (
	"fmt"

	"example.com/commands"
	"example.com/logMonitor"
)

func main() {
	logMonitor.ShowInfo()

	ipList := logMonitor.MonitLogs()
	badActors := logMonitor.GetMalIps(ipList)
	fmt.Println("Dropping malicious hosts...")
	commands.DropBadActors(badActors)
}
