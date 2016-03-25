package main

import "fmt"
import "./config"

func main() {
    fmt.Println(config.FindConfigFile())
}
