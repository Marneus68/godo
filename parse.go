package main

import (
	"flag"
	"fmt"
	"github.com/Marneus68/godo/config"
	"github.com/Marneus68/godo/job"
	"github.com/Marneus68/godo/servers"
	"github.com/Marneus68/godo/starter"
	"log"
	"os"
	"strings"
)

type OptionsFunc func(ex string, args []string)

var f = flag.NewFlagSet("standard flags", flag.ContinueOnError)

var defConfigFilePath *string

var defName *string
var defType *string
var defPort *string
var defWeb *bool
var defWebPort *string
var defTags *string
var defComm *string

// Defines the possible optional flags for most commands
// (create, config, start, restart)
func SetStandardFlags() {
	defConfigFilePath = f.String("conf", "", "Path to the config file")
	defName = f.String("name", "", "The name of the instance")
	defType = f.String("type", "", "The type of the instance")
	defPort = f.String("port", "", "Port used for communication")
	defWeb = f.Bool("web", false, "Enables the web interface")
	defWebPort = f.String("webport", "", "Port of the web interface")
	defTags = f.String("tags", "", "Comma separated list of tags")
}

// Parse the standard flags and returns a config
func ParseStandardFlags(args []string, t config.InstanceType) (ret config.Config) {
	SetStandardFlags()
	ret.Type = t
	if len(args) > 0 {
		if err := f.Parse(args); err != nil {
			log.Fatal("Error while parsing command line parameters")
		}
		if f.Parsed() {
			ret.Type = config.Master
			if *defName != "" {
				ret.Name = *defName
			}
			if *defType != "" {
				switch strings.TrimSpace(strings.ToLower(*defType)) {
				case "master":
					ret.Type = config.Master
				case "servant":
					ret.Type = config.Servant
				case "slave":
					ret.Type = config.Slave
				}
			}
			if *defPort != "" {
				ret.Port = *defPort
			}
			ret.Web = *defWeb
			if *defWebPort != "" {
				ret.WebPort = *defWebPort
			}
		}
	}
	return ret
}

// Defines the possible flags for the Job control mode
func SetJobFlags() {
	defName = f.String("name", "", "Name of the job")
	defComm = f.String("command", "", "Command for the job")
	defTags = f.String("tags", "", "Tags for the job")
}

// Parse the job flags
func ParseJobFlags(args []string) (ret job.Job) {
	// TODO
	return ret
}

var Options = map[string]OptionsFunc{
	"create": func(ex string, args []string) {
		switch {
		case len(args) > 1:
			var t config.InstanceType
			switch strings.ToLower(args[1]) {
			case "master":
				t = config.Master
			case "servant":
				t = config.Servant
			case "slave":
				t = config.Slave
			default:
				log.Fatal("Unknown instance type  \"", args[1], "\"... Aborting.")
			}
			fmt.Println("Creating a", args[1], "godo instance")
			c := ParseStandardFlags(args[2:], t)
			fmt.Println(c.ToString())
		default:
			FlagError()
		}
	},
	"config": func(ex string, args []string) {
		switch {
		case len(args) > 1:
			switch strings.ToLower(args[1]) {
			case "edit":
				fmt.Println("Editing local configuration file")
				if len(args) > 2 {
					// TODO: Read the local config file
					// Parse the command line config
					c := ParseStandardFlags(args[2:], config.Master)
					// TODO: Merge the command line config to the local config
					// TODO: Write the result
					fmt.Println(c.ToString())
				} else {
					fmt.Println("No parameters provided...")
					fmt.Println("Attempting to open config file with standard text editor...")
					os.Exit(1)
				}
			default:
				FlagError()
			}
		default:
			FlagError()
		}
	},
	"start": func(ex string, args []string) {
		fmt.Println("start")
		c := config.NewConfig()
		starter.Start(ex, args, *c)
	},
	"restart": func(ex string, args []string) {
		fmt.Println("restart")
	},
	"stop": func(ex string, args []string) {
		fmt.Println("stop")
		c := config.NewConfig()
		starter.Stop(ex, args, *c)
	},
	"status": func(ex string, args []string) {
		fmt.Println("status")
	},
	"job": func(ex string, args []string) {
		SetJobFlags()
		switch {
		case len(args) > 1:
			switch strings.ToLower(args[1]) {
			case "create":
				fmt.Println("CREATING NEW JOB")
			case "delete":
				fmt.Println("DELETING JOB")
			case "start":
				fmt.Println("STARTING JOB")
			case "edit":
				fmt.Println("EDITING JOB")
			default:
				FlagError()
			}
		default:
			FlagError()
		}
	},
	"deamon": func(ex string, arg []string) {
		fmt.Println("Starting the godo demon...")
		// TODO: Read config from config file
		c := config.NewConfig()
		// TODO: Create config from command line arguments

		// TODO: Merge command line config to config from file

		servers.Start(*c)
	},
}

func FlagError() {
	fmt.Println("Not enough parameters provided for this command.")
	PrintUsage()
	os.Exit(1)
}
