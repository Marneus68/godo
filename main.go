package main

import (
	"./control"
	"./usage"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	switch {
	case len(args) == 0:
		usage.Print()
		os.Exit(1)
	case len(args) > 0:
		if fun, ok := control.Map[strings.ToLower(args[0])]; ok {
			fun()
		} else {
			usage.Print()
		}
	}
}
