package main

import (
	"fmt"
	"log"
	"os"

	"example.com/commands"
	"example.com/logMonitor"
	"example.com/sha256checksums"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Overseer"
	app.Usage = "Log monitoring and intrusion prevention tool."

	app.Commands = []cli.Command{
		{
			Name:  "full",
			Usage: "Full scan of auth.log file and drop any found malicious hosts.",
			Action: func(*cli.Context) {
				logMonitor.ShowInfo()
				ipList := logMonitor.MonitLogs()
				badActors := logMonitor.GetMalIps(ipList)
				maldns := logMonitor.MalDns()
				for i := range maldns {
					fmt.Println("Malicious DNS found: ", i+1, maldns[i])
				}
				fmt.Println("Dropping malicious hosts...")
				commands.DropBadActors(badActors)
				commands.DropBadActors(maldns)
			},
		},
		{
			Name:  "checksums",
			Usage: "Generate sha-256 hash of all files in /bin for file manipulation checks",
			Action: func(*cli.Context) {
				sums := sha256checksums.GenSums()
				sha256checksums.GenSumFile(sums)
				sha256checksums.CheckSums()
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
