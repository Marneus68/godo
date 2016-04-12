package server

import (
	"fmt"
	"github.com/Marneus68/godo/config"
	"log"
	"os"
	"strconv"
	"strings"
)

// Start the godo server and writes the pid of the newly created instance
// if pidfile isn't an empty string
func Start(con config.Config, pidfile string) {
	if strings.Compare(pidfile, "") != 0 {
		f, err := os.Create(pidfile)
		if err != nil {
			log.Fatal("This shouldn't be happening!")
		}
		_, err = f.WriteString(strconv.Itoa(os.Getpid()))
		if err != nil {
			log.Fatal("This shouldn't be happening!")
		}
	}
	fmt.Println("Starting godo with pid " + strconv.Itoa(os.Getpid()) + "...")
}
