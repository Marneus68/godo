package main

import (
	//"./control"
	//"./usage"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	switch {
	case len(args) == 0:
		PrintUsage()
		os.Exit(1)
	case len(args) > 0:
		if fun, ok := Options[strings.ToLower(args[0])]; ok {
			fun(args)
		} else {
			PrintUsage()
		}
	}
}
