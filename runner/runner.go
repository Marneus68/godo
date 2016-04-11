package runner

import (
	"fmt"
	"github.com/Marneus68/godo/config"
	"log"
	"os"
	"path/filepath"
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
	f, err := os.Open(pidfile)
	if err != nil {
		log.Fatal("Error opening \"" + pidfile + "\"")

		// If file doesn't exist, we can start the server which will write it

		// If you don't have the permission to real the file... huh
	}
}

func Restart(con config.Config) {

}

func Stop(con config.Config) {

}
