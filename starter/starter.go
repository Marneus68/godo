// Package handling the start, restart and stop operations
package starter

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Marneus68/godo/config"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// Name of the file used to store the pid of the current instance
const PIDFILE_NAME string = "pidfile"

// TODO: Create a function (CheckPidfile) that either returns the pid or an error

// Returns the content of the pidfile or an error
func ReadPidfile() (pid int, pidfile string, err error) {
	pidfile = filepath.Join(config.ConfigDirectory(), PIDFILE_NAME)
	_, err = os.Open(pidfile)
	if err != nil {
		//log.Fatal("Error opening \"" + pidfile + "\"")
		fmt.Println("Error opening pidfile (" + pidfile + ")...")
		if os.IsPermission(err) {
			log.Fatal("Permission errors while trying to open the pidfile (" + pidfile + ")")
		}
		if os.IsNotExist(err) {
			//server.Start(con, pidfile)
			return 0, pidfile, errors.New("The pidfile doesn't exist (" + pidfile + ")")
		}
	}

	f, err := os.Open(pidfile)
	if err != nil {
		log.Fatal("Unkown error.")
	}

	// Read the pid
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		pid, err = strconv.Atoi(line)
		if err != nil {
			log.Fatal("Error reading pidfile (" + pidfile + ")")
		}
		break
	}
	return pid, pidfile, nil
}

// Start godo with the deamon argument
func startDeamon(ex string, args []string, con config.Config) {
	arr := []string{"daemon"}
	cmd := exec.Command(ex, arr...)
	cmd.Start()
}

// Start the server is no other instance is running
//
// Check if an instance of godo is already running, if no other
// instance have been found, we start the server which will take
// care of the whole forking thing
func Start(ex string, args []string, con config.Config) {
	//p, err := os.FindProcess(
	fmt.Println("[CONFIG PATH]", config.ConfigDirectory())
	fmt.Println("[CONFIG FILE]", config.ConfigFile())
	fmt.Println("[JOBS FILE PATH]", config.JobsDirectory())
	fmt.Println("[PIDFILE]: ", filepath.Join(config.ConfigDirectory(), PIDFILE_NAME))

	pid, pidfile, err := ReadPidfile()

	if pid == 0 {
		fmt.Println("Invalid pid... Creating new godo instance.")
		os.Remove(pidfile)
		//server.Start(con, pidfile)
	}
	//fmt.Println("pid in pidfile : " + strconv.Itoa(pid) + " in (" + pidfile + ")")

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
		//server.Start(con, pidfile)
	}
}

// Restart the server is one is already running
func Restart(ex string, args []string, con config.Config) {

}

// Stop the server if one is running
func Stop(ex string, args []string, con config.Config) {

}
