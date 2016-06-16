// Package server implements the godo server and web server
package servers

import (
	"fmt"
	"github.com/Marneus68/godo/config"
	"github.com/Marneus68/godo/static"
	//"github.com/Marneus68/godo/job"
	"github.com/Marneus68/godo/starter"
	//"html"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Server static resources
//var resources *map[string][]byte

// Normalize the port string
func NormalizePortString(port string) string {
	return ":" + strings.TrimPrefix(port, ":")
}

// Start the godo server and writes the pid of the newly created instance
// if pidfile isn't an empty string
func Start(con config.Config) {
	// Set the global static resources of the web server
	//resources = res

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
	WebServer(con)
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
	// Add handlers for all the static content
	for k, v := range static.Content {
		http.HandleFunc(strings.Replace(k, "static", "", 1), func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, string(v))
		})
	}
	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the end install gentoo")
	})
	// Serve the static content
	http.ListenAndServe(":5678", nil)
}

func handleIncomming(conn net.Conn) {

}
