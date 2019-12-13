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

var help, version bool

func main() {
	flag.Usage = myUsage
	flag.BoolVar(&help, "help", false, "show help")
	flag.BoolVar(&version, "version", false, "show version")
	flag.Parse()

	cmd := flag.Arg(0)
	if version {
		fmt.Printf("remo version %s\n", remo.GetCurrentVersion())
		os.Exit(0)
	} else if help || cmd == "" {
		flag.Usage()
		os.Exit(0)
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
	case "start", "ohayo":
		err = cli.StartRemoteWork()
	case "pause":
		err = cli.PauseRemoteWork()
	case "resume", "unpause":
		err = cli.ResumeRemoteWork()
	case "move-office":
		err = cli.MoveToOffice()
	case "arrive-office":
		err = cli.ArriveAtOffice()
	case "done", "otsu":
		err = cli.FinishRemoteWork()
	case "init":
		_, err = remo.CreateConfigFile()
	default:
		flag.Usage()
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
