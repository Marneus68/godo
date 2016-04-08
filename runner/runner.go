package runner

import (
	"fmt"
	"github.com/Marneus68/godo/config"
)

// Check if an instance of godo is already running, if no other
// instance have been found, a new one is made
func Start(con config.Config) {
	//p, err := os.FindProcess(
	fmt.Println("[CONFIG DIRECTORY]: " + config.ConfigDirectory())
}

func Restart(con config.Config) {

}

func Stop(con config.Config) {

}
