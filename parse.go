package main

import (
	"flag"
	"fmt"
	"github.com/Marneus68/godo/config"
	"log"
	"os"
)

type OptionsFunc func(args []string)

var f = flag.NewFlagSet("standard flags", flag.ContinueOnError)

var defName *string
var defPort *string
var defWeb *bool
var defWebPort *string

// Defines the possible optional flags for most commands (create, config)
func SetStandardFlags() {
	defName = f.String("name", "", "The name of the instance")
	defPort = f.String("port", "", "Port used for communication")
	defWeb = f.Bool("web", false, "Enables the web interface")
	defWebPort = f.String("webport", "", "Port of the web interface")
}

// Defines the possible flags for the Job control mode
func SetJobFlags() {

}

var Options = map[string]OptionsFunc{
	"create": func(args []string) {
		SetStandardFlags()
		switch {
		case len(args) > 1:
			c := config.NewConfig()
			switch args[1] {
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
		case len(args) == 1:
			fallthrough
		default:
			fmt.Println("Not enough parameters provided for this command.")
			PrintUsage()
			os.Exit(1)
		}
	},
	"config": func(args []string) {
		fmt.Println("config")
	},
	"start": func(args []string) {
		fmt.Println("start")
	},
	"restart": func(args []string) {
		fmt.Println("restart")
	},
	"stop": func(args []string) {
		fmt.Println("stop")
	},
	"status": func(args []string) {
		fmt.Println("status")
	},
	"job": func(args []string) {
		fmt.Println("job")
	},
}
