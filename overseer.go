package main

import (
	"fmt"

	"example.com/commands"
	"example.com/logMonitor"
	"example.com/sha256checksums"
)

func main() {
	logMonitor.ShowInfo()

	ipList := logMonitor.MonitLogs()
	badActors := logMonitor.GetMalIps(ipList)
	maldns := logMonitor.MalDns()
	for i := range maldns {
		fmt.Println("Malicious Rdns ", i+1, maldns[i])
	}
	fmt.Println("Dropping malicious hosts...")
	sums := sha256checksums.GenSums()
	sha256checksums.GenSumFile(sums)
	sha256checksums.CheckSums()
	commands.DropBadActors(badActors)
	commands.DropBadActors(maldns)
}
