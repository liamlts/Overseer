package main

import (
	"fmt"
	"os/exec"
)

func DropBadActors(mIps []string) {
	firewall := "ufw"

	arg := "deny"
	arg1 := "from"
	var curIP string

	for index := range mIps {
		curIP = mIps[index]
		cmd := exec.Command(firewall, arg, arg1, curIP)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(curIP, " ", string(stdout))
	}
}
