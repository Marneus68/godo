// Package responsible for handling the networking
package server

import (
	"fmt"
	"github.com/Marneus68/godo/config"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func NormalizePortString(port string) string {
	return ":" + strings.TrimPrefix(port, ":")
}

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
		if con.Web {
			go WebServer(con)
		}
		break
	}
	go IncommingServer(con)
}

func IncommingServer(con config.Config) {
	ln, err := net.Listen("tcp", NormalizePortString(con.Port))
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleIncomming(conn)
	}
}

func WebServer(con config.Config) {
	http.ListenAndServ(NormalizePortString(con.WebPort), nil)
}

func handleIncomming(conn net.Conn) {

}
