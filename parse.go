package main

import (
	"flag"
	"fmt"
	"github.com/Marneus68/godo/config"
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
	defPort = f.String("port", "", "Port used for communication")
	defWeb = f.Bool("web", false, "Enables the web interface")
	defWebPort = f.String("webport", "", "Port of the web interface")
	defTags = f.String("tags", "", "Comma separated list of tags")
}

// Defines the possible flags for the Job control mode
func SetJobFlags() {
	defName = f.String("name", "", "Name of the job")
	defComm = f.String("command", "", "Command for the job")
	defTags = f.String("tags", "", "Tags for the job")
}

var Options = map[string]OptionsFunc{
	"create": func(ex string, args []string) {
		SetStandardFlags()
		switch {
		case len(args) > 1:
			c := config.NewConfig()
			switch strings.ToLower(args[1]) {
			case "master":
				c.Type = config.Master
			case "servant":
				c.Type = config.Servant
			case "slave":
				c.Type = config.Slave
			default:
				log.Fatal("Unknown instance type  \"", args[1], "\"... Aborting.")
			}
			fmt.Println("Creating a", args[1], "godo instance")
			if len(args) > 2 {
				if err := f.Parse(args[2:]); err != nil {
					log.Fatal("Error while parsing command line parameters")
				}
				if f.Parsed() {
					c.Type = config.Master
					if *defName != "" {
						c.Name = *defName
					}
					if *defPort != "" {
						c.Port = *defPort
					}
					c.Web = *defWeb
					if *defWebPort != "" {
						c.WebPort = *defWebPort
					}
				}
			}
			fmt.Println(c.ToString())
		default:
			FlagError()
		}
	},
	"config": func(ex string, args []string) {
		SetStandardFlags()
		switch {
		case len(args) > 1:
			c := config.NewConfig()
			switch strings.ToLower(args[1]) {
			case "edit":
				if len(args) > 2 {
					if err := f.Parse(args[2:]); err != nil {
						log.Fatal("Error while parsing command line parameters")
					}
					if f.Parsed() {
						c.Type = config.Master
						if *defName != "" {
							c.Name = *defName
						}
						if *defPort != "" {
							c.Port = *defPort
						}
						c.Web = *defWeb
						if *defWebPort != "" {
							c.WebPort = *defWebPort
						}
					}
				} else {
					fmt.Println("No parameters provided...")
					fmt.Println("Attempting to open config file with standard text editor...")
					os.Exit(1)
				}
			default:
				FlagError()
			}
			fmt.Println(c.ToString())
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
	"demon": func(ex string, arg []string) {
		fmt.Println("Starting the godo demon...")
		c := config.NewConfig()
		servers.Start(*c)
	},
}

func FlagError() {
	fmt.Println("Not enough parameters provided for this command.")
	PrintUsage()
	os.Exit(1)
}
