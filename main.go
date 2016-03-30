package main

import (
	"./control"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	switch {
	case len(args) == 0:
		fmt.Print("usage\n")
		os.Exit(1)
	case len(args) > 0:
		control.Map[strings.ToLower(args[0])](args)
		/*
		   switch strings.ToLower(args[0]) {
		   case "create":
		       control.Create()
		   case "config":
		       control.Config()
		   case "start":
		       control.Start()
		   case "restart":
		       control.Restart()
		   case "stop":
		       control.Stop()
		   case "job":
		       control.Job()
		   }
		*/
	}
}
