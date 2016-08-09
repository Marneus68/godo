// Package server implements the godo server and web server
package servers

import (
	"fmt"
	"github.com/Marneus68/godo/config"
	"github.com/Marneus68/godo/starter"
	"github.com/Marneus68/godo/static"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Local server configuration
var localConfig config.Config

// Normalize the port string
func NormalizePortString(port string) string {
	return ":" + strings.TrimPrefix(port, ":")
}

// Start the godo server and writes the pid of the newly created instance
// if pidfile isn't an empty string
func Start(con config.Config) {
	// Set the local configuration for this instance of the server
	localConfig = con

	// Find the pidfile
	_, pidfile, _ := starter.ReadPidfile()

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

	/*
		switch con.Type {
		case config.Master:
			if con.Web {
				go WebServer(con)
			}
			break
		}
		go IncommingServer(con)

		for true {
		}
	*/
	WebServer()
}

func IncommingServer() {
	ln, err := net.Listen("tcp", ":"+localConfig.Port)
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

func WebServer() {
	// Add handlers for all the static content
	for index, content := range static.Content {
		if !strings.Contains(index, "index.html") {
			http.HandleFunc(strings.Replace(index, "www/", "", -1), func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, string(content))
			})
		}
	}
	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(static.Content["www/index.html"]))
	})

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, string(static.Content["www/index.html"]))
		})
	*/

	// Serve the static content
	http.ListenAndServe(":5678", nil)
}

func handleIncomming(conn net.Conn) {

}
