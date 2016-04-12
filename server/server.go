package server

import (
	"fmt"
	"github.com/Marneus68/godo/config"
	"log"
	"os"
	"strconv"
)

func Start(con config.Config, pidfile string) {
	f, err := os.Create(pidfile)
	if err != nil {
		log.Fatal("This shouldn't be happening!")
	}
	fmt.Println("Starting godo with pid " + strconv.Itoa(os.Getpid()) + "...")
	_, err = f.WriteString(strconv.Itoa(os.Getpid()))
	if err != nil {
		log.Fatal("This shouldn't be happening!")
	}
}
