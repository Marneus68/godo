package main

import (
	"flag"
	"fmt"
)

type OptionsFunc func(args []string)

var Options = map[string]OptionsFunc{
	"create":  func([]string) { fmt.Println("create") },
	"config":  func([]string) { fmt.Println("config") },
	"start":   func([]string) { fmt.Println("start") },
	"restart": func([]string) { fmt.Println("restart") },
	"stop":    func([]string) { fmt.Println("stop") },
	"status":  func([]string) { fmt.Println("status") },
	"job":     func([]string) { fmt.Println("job") },
}
