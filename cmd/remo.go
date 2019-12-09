package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mactkg/remo"
)

func myUsage() {
	// TODO: refactor to struct `Action`
	fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
	fmt.Printf("\nOptions:\n")
	flag.PrintDefaults()

	fmt.Printf("Actions:\n")
	fmt.Printf("  init:\n\tCreate config file\n")
	fmt.Printf("  start,ohayo:\n\tStart remote working\n")
	fmt.Printf("  pause:\n\tPause working\n")
	fmt.Printf("  resume,unpause:\n\tResume working\n")
	fmt.Printf("  arrive-office:\n\tNotify you arrived at office and resuming work\n")
	fmt.Printf("  move-office:\n\tStop remote work and move to office\n")
	fmt.Printf("  finish,done,otsu:\n\tFinish working\n")
}

func loadConfig() (*remo.Config, error) {
	config := &remo.Config{}

	configPath, err := remo.GetDefaultConfigPath()
	if err != nil {
		return nil, err
	}
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	decoder := remo.NewDecoder(configFile)
	decoder.Decode(config)

	return config, nil
}

func main() {
	flag.Usage = myUsage
	flag.Parse()

	cmd := flag.Arg(0)
	help := flag.Lookup("h")
	if help != nil || cmd == "" {
		flag.Usage()
		return
	}

	if cmd == "init" {
		f, err := remo.CreateConfigFile()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("created new config file to %v", f.Name())
		os.Exit(0)
	}

	config, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't load config:\n%v\n", err)
		fmt.Fprintf(os.Stderr, "hint: you can create config run with 'remo init'\n")
		os.Exit(1)
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
	case "arrive-office":
		cli.ArriveAtOffice()
	case "done":
	case "otsu":
		cli.FinishRemoteWork()
	case "init":
		remo.CreateConfigFile()
	default:
		flag.Usage()
	}
}
