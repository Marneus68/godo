// Package doing all the heavy lifting, opening the various sockets,
// listening or sending messages and starting the appropriate goroutines
// when needed
package server

import (
	"fmt"
	"github.com/Marneus68/godo/config"
	"log"
	"net"
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

	switch con.Type {
	case config.Master:
		break
	case config.Slave:
		break
	case config.Servant:
		break
	}
}

func IncommingListener(con config.Config) {
	ln, err := net.Listen("tcp", ":"+strings.TrimPrefix(con.Port, ":"))
	if err != nil {

	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleIncomming(conn)
	}
}

func handleIncomming(conn net.Conn) {

}
