package main

import (
    "os"
    "fmt"
)

func main() {
	//fmt.Println(config.ConfigFile())
    args := os.Args[1:]

    switch {
    case len(args) == 0:
        fmt.Print("usage")
        os.Exit(1)
    case len(args) > 1:
        switch os.Args[1] {
        case "create":

        case "config":

        case "start":

        case "restart":

        case "strop":

        case "job":

        }
    }
}
