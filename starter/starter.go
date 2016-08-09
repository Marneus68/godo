// Package starter implements the main functions called after the argument parsing
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

// TODO: make the pid variable a package variable

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

// Starts godo with the deamon argument
func StartDeamon(ex string, args []string, con config.Config) {
	fmt.Println("Starting godo in deamon mode...")
	arr := []string{"deamon"}
	cmd := exec.Command(ex, arr...)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
}

// Kill the daemon process
func KillDaemon(pid int) {
	p, err := os.FindProcess(pid)
	if err == nil {
		err = p.Kill()
		if err != nil {
			log.Fatal("Error while attempting to kill the already running godo instance")
		}
	}
}

type ActionFunc func(pid int, ex string, args []string, con config.Config)

// Checks is the pidfile exists, if it exists we check if the an associated process exists
// If an instance of godo is found, `doIfExist` is run, if no instance is found,
// `doIfDoesntExist` is run instead
func checkForGodoInstance(ex string, args []string, con config.Config, doIfExist ActionFunc, doIfDoesntExist ActionFunc) {
	pid, pidfile, err := ReadPidfile()
	if pid == 0 {
		// the pidfile doesn't exist
		fmt.Println("Invalid pid... Creating new godo instance.")
		os.Remove(pidfile)
		doIfDoesntExist(pid, ex, args, con)
		return
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
		doIfExist(pid, ex, args, con)
	} else {
		// the pid doesn't represent a running instance, we delete the pidfile
		os.Remove(pidfile)
		doIfDoesntExist(pid, ex, args, con)
		return
	}
}

// Start the server is no other instance is running
//
// Check if an instance of godo is already running, if no other
// instance have been found, we start the server which will take
// care of the whole forking thing
func Start(ex string, args []string, con config.Config) {
	fmt.Println("[CONFIG PATH]", config.ConfigDirectory())
	fmt.Println("[CONFIG FILE]", config.ConfigFile())
	fmt.Println("[JOBS FILE PATH]", config.JobsDirectory())
	fmt.Println("[PIDFILE]: ", filepath.Join(config.ConfigDirectory(), PIDFILE_NAME))

	checkForGodoInstance(ex, args, con,
		func(pid int, ex string, args []string, con config.Config) {
			log.Fatal("An instance of godo is already running!")
		},
		func(pid int, ex string, args []string, con config.Config) {
			StartDeamon(ex, args, con)
		},
	)
}

// Restart the server is one is already running
func Restart(ex string, args []string, con config.Config) {
	checkForGodoInstance(ex, args, con,
		func(pid int, ex string, args []string, con config.Config) {
			KillDaemon(pid)
			StartDeamon(ex, args, con)
		},
		func(pid int, ex string, args []string, con config.Config) {
			StartDeamon(ex, args, con)
		},
	)
}

// Stop the server if one is running
func Stop(ex string, args []string, con config.Config) {
	checkForGodoInstance(ex, args, con,
		func(pid int, ex string, args []string, con config.Config) {
			KillDaemon(pid)
		},
		func(pid int, ex string, args []string, con config.Config) {
			log.Fatal("No godo instance found!")
		},
	)
}
