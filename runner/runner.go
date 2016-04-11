package runner

import (
	"bufio"
	"fmt"
	"github.com/Marneus68/godo/config"
	"github.com/Marneus68/godo/server"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

// Name of the file used to store the pid of the current instance
const PIDFILE_NAME string = "pidfile"

// Check if an instance of godo is already running, if no other
// instance have been found, a new one is made
func Start(con config.Config) {
	//p, err := os.FindProcess(
	fmt.Println("[CONFIG PATH]", config.ConfigDirectory())
	fmt.Println("[CONFIG FILE]", config.ConfigFile())
	fmt.Println("[JOBS FILE PATH]", config.JobsDirectory())
	fmt.Println("[PIDFILE]: ", filepath.Join(config.ConfigDirectory(), PIDFILE_NAME))

	pidfile := filepath.Join(config.ConfigDirectory(), PIDFILE_NAME)
	_, err := os.Open(pidfile)
	if err != nil {
		//log.Fatal("Error opening \"" + pidfile + "\"")
		fmt.Println("Error opening pidfile (" + pidfile + ")...")
		if os.IsPermission(err) {
			log.Fatal("Permission errors.")
		}
		if os.IsNotExist(err) {
			server.Start(con)
			return
		}
	}

	f, err := os.Open(pidfile)
	if err != nil {
		log.Fatal("Unkown error.")
	}

	// Read the pid
	var pid int = 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		pid, err = strconv.Atoi(line)
		if err != nil {
			log.Fatal("Error reading pidfile (" + pidfile + ")")
		}
		break
	}
	if pid == 0 {
		fmt.Println("Invalid pid... Creating new godo instance.")
		os.Remove(pidfile)
		server.Start(con)
	}
	fmt.Println("pid in pidfile : " + strconv.Itoa(pid) + " in (" + pidfile + ")")

	var running bool = false

	p, err := os.FindProcess(pid)
	if err == nil {
		err := p.Signal(syscall.Signal(0))
		if err == nil {
			running = true
		}
	}

	if running {
		// godo is already running !
		log.Fatal("An instance of godo is already running")
	} else {
		// the pid doesn't represent a running instance, we delete the pidfile
		os.Remove(pidfile)
		server.Start(con)
	}
}

func Restart(con config.Config) {

}

func Stop(con config.Config) {

}
