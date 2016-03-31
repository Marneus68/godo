package config

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
	"strings"
)

// Enumeration describing the instance type, cen be either master, slave
// or servant (both)
type InstanceType int

const (
	Master InstanceType = iota
	Slave
	Servant
)

// Structure describing the godo instance configuration
//
// Name : instance name
// Type : instance type
// Os : os of the instance
// Port : port godo listens on if it's either a slave or a servant
// Tags : job tags accepted by this instance if it's either a slave or servant
// Web : boolean, tells if the instance has its web interface enabled
// WebPort : port for the web interface
type Config struct {
	Name    string
	Type    InstanceType
	Os      string
	Port    string
	Tags    []string
	Web     bool
	WebPort string
	Slaves  []string
}

var locations = map[string]string{
	"/godo/godo.conf":                 "/godo/jobs.d",
	"/etc/godo/godo.conf":             "/etc/godo/jobs.d",
	"/usr/local/share/godo/godo.conf": "/usr/local/share/godo/jobs.d",
	"~/.config/godo/godo.conf":        "~/.config/godo/jobs.d/",
	"~/godo/godo.conf":                "~/godo/jobs.d/",
}

var configFile = ""
var jobDirectory = ""

// Looks for the configuration file from multiple possible locations
// returns the full absolute path to the first configuration file found
func ConfigFile() string {
	u, _ := user.Current()
	homeDir := u.HomeDir

	if homeDir == "" {
		log.Fatal("Unable to find the home directory of the current user.")
	}

	homeDir = fmt.Sprint(homeDir + "/")

	var i = 0
	for k, _ := range locations {
		path := k
		if path[:2] == "~/" {
			path = strings.Replace(path, "~/", homeDir, 1)
		}

		fmt.Printf("Checking for file \"%s\"\n", path)
		_, err := os.Stat(path)
		if err == nil {
			fmt.Print("FOUND")
			configFile = k
			break
		}
		i++
	}
	if configFile == "" {
		log.Fatal("Could not find configuration file.")
	}
	return configFile
}

// Looks for the jobs directory
// returns the full absolute path to the jobs directory corresponding to the
// config file used
func JobDirectory() string {
	return ""
}

// Read configuration from a config file
func (config Config) FromConfigFile() {

}

// Read configuration from a command line parameters
func (config Config) FromArgs() {

}

// Constructor for the config struct
func NewConfig() *Config {
	c := new(Config)
	name, err := os.Hostname()
	if err == nil {
		c.Name = name
	} else {
		c.Name = ""
	}
	c.Type = Master
	c.Os = runtime.GOOS
	c.Port = "8008"
	c.Tags = make([]string, 0)
	c.Web = false
	c.WebPort = "8888"
	c.Slaves = make([]string, 0)
	return c
}

// Prints a human readable rundown of the configuration
func (config Config) ToString() string {
	ret := ""
	switch config.Type {
	case Master:
		ret = fmt.Sprintf(
			"[master godo instance]\nName : %s\nOS : %s", config.Name, config.Os)

	case Servant:
	case Slave:
	}
	return ret
}
