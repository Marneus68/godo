package main

import (
	"flag"
	"fmt"
	"github.com/Marneus68/godo/config"
	"log"
	"os"
)

type OptionsFunc func(args []string)

var f = flag.NewFlagSet("flags", flag.ContinueOnError)

// Default values set by the flag parser
var defName *string
var defHost *string
var defPort *string
var defOs *string
var defWeb *bool
var defWebPort *string

// Defines the possible flags for most modes
func SetStandardFlags() {
	defName = f.String("name", "godo_instance", "The name of the instance")
	defHost = f.String("host", "localhost", "The host name (or IP)")
	defPort = f.String("port", "8008", "Port used for communication")
	defOs = f.String("os", "linux", "OS the godo instance is running on")
	defWeb = f.Bool("web", false, "Enables the web interface")
	defWebPort = f.String("webport", "8888", "Port of the web interface")
}

// Defins the possible flags for the Job control mode
func SetJobFlags() {
}

var Options = map[string]OptionsFunc{
	"create": func(args []string) {
		switch {
		case len(args) == 1:
			fmt.Println("Creating godo master instance with default parameters:")
		case len(args) > 1:
			fmt.Println("Creating gogo master instance with custom parameters:")
			if err := f.Parse(args[2:]); err != nil {
				log.Fatal("Error while parsing command line parameters")
			}
			c := config.NewConfig()
			if f.Parsed() {
				c.Type = config.Master
				if defName != nil {
					c.Name = *defName
				}
				if defOs != nil {
					c.Os = *defOs
				}
				//c.Host = *defHost
				if defPort != nil {
					c.Port = *defPort
				}
				if defWeb != nil {
					c.Web = *defWeb
				}
				if defWebPort != nil {
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
