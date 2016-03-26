package main

import (
	"fmt"
)

var Controls = map[string]func(){
	"create":  func() { fmt.Println("create") },
	"config":  func() { fmt.Println("config") },
	"start":   func() { fmt.Println("start") },
	"restart": func() { fmt.Println("restart") },
	"stop":    func() { fmt.Println("stop") },
	"status":  func() { fmt.Println("status") },
	"job":     func() { fmt.Println("job") },
}
