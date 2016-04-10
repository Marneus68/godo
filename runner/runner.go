package runner

import (
	"fmt"
	"github.com/Marneus68/godo/config"
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
	fmt.Println("[JOBS FILE NAME]", config.JobsFile())
	fmt.Println("[JOBS FILE PATH]", config.JobsDirectory())
	fmt.Println("[PIDFILE]: ", filepath.Join(config.ConfigDirectory(), PIDFILE_NAME))
	// if the pidfile doesn't exist we create our instance and write it

	// if pidfile exists we read it
	// we check if there is a process with the same pid and we attempt to kill it

	// we create our instance, write the pid in the pidfile and fork out
}

func Restart(con config.Config) {

}

func Stop(con config.Config) {

}
