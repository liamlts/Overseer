package main

import (
	"fmt"
	"os/exec"

	"example.com/logMonitor"
)

func DropBadActors(mIps []string) {
	firewall := "ufw"

	arg := "deny"
	arg1 := "from"
	var curIP string

	for index, _ := range mIps {
		curIP = mIps[index]
		cmd := exec.Command(firewall, arg, arg1, curIP)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(curIP, " ", string(stdout))
	}
}

func main() {
	logMonitor.ShowInfo()

	ipList := logMonitor.MonitLogs()
	badActors := logMonitor.GetMalIps(ipList)
	fmt.Println("Dropping malicious hosts...")
	DropBadActors(badActors)
}
