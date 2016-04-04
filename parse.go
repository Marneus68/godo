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

// Default values set by the flag parser
var defName *string
var defPort *string
var defWeb *bool
var defWebPort *string

// Defines the possible flags for most modes
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
		case len(args) == 1:
			fmt.Println("Creating godo master instance with default parameters:")
		case len(args) > 1:
			if err := f.Parse(args[1:]); err != nil {
				log.Fatal("Error while parsing command line parameters")
			}
			fmt.Println("Creating gogo master instance with custom parameters:")
			c := config.NewConfig()
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
			fmt.Println(c.ToString())
		default:
			fmt.Println("Not enough parameters provided for this command.\n")
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
