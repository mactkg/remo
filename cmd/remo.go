package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mactkg/remo"
)

var mainPostChannel string

func myUsage() {
	// TODO: refactor to struct `Action`
	fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
	fmt.Printf("\nOptions:\n")
	flag.PrintDefaults()

	fmt.Printf("Actions:\n")
	fmt.Printf("  start,ohayo:\n\tStart remote working\n")
	fmt.Printf("  pause:\n\tPause working\n")
	fmt.Printf("  resume,unpause:\n\tResume working\n")
	fmt.Printf("  move-office:\n\tStop remote work and move to office\n")
	fmt.Printf("  finish,done,otsu:\n\tFinish working\n")
}

func main() {
	flag.Usage = myUsage
	flag.StringVar(&mainPostChannel, "mainPostCh", "", "where to post")
	flag.Parse()

	cmd := flag.Arg(0)
	help := flag.Lookup("h")
	if help != nil || cmd == "" {
		flag.Usage()
		return
	}

	token := os.Getenv("REMO_SLACK_TOKEN")
	config := remo.Config{
		SlackToken:      token,
		MainPostChannel: mainPostChannel,
	}

	cli := remo.New(config)

	switch cmd {
	case "start":
	case "ohayo":
		cli.StartRemoteWork()
	case "pause":
		cli.PauseRemoteWork()
	case "resume":
	case "unpause":
		cli.ResumeRemoteWork()
	case "move-office":
		cli.MoveToOffice()
	case "done":
	case "otsu":
		cli.FinishRemoteWork()
	}
}
