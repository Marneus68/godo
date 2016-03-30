package control

import (
	"fmt"
)

var Map = map[string]func(){
	"create":  func() { fmt.Println("create") },
	"config":  func() { fmt.Println("config") },
	"start":   func() { fmt.Println("start") },
	"restart": func() { fmt.Println("restart") },
	"stop":    func() { fmt.Println("stop") },
	"status":  func() { fmt.Println("status") },
	"job":     func() { fmt.Println("job") },
}

/*
func Create(args []string) {
    fmt.Println("CREATE")
}

func Config(args []string) {
    fmt.Println("CONFIG")
}

func Start(args []string) {
    fmt.Println("START")
}

func Restart(args []string) {
    fmt.Println("RESTART")
}

func Stop(args []string) {
    fmt.Println("STOP")
}

func Status(args []string) {
    fmt.Println("STATUS")
}

func Job(args []string) {
    fmt.Println("JOB")
}
*/
