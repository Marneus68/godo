package config

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"reflect"
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

// Enumeration describing the way slaves or servants are selected to submit a job
type SlaveSelectMode int

const (
	// Round robin between all slaves
	RoundRobin SlaveSelectMode = iota
	// Round robin between all slaves and the current instance itself if it's a servant
	RoundRobinSelf
)

// Default tupe
const DEFAULT_TYPE InstanceType = Master

// Default name
var DEFAULT_NAME string = ""

// Default port
const DEFAULT_PORT string = "8008"

// Default tags
var DEFAULT_TAGS []string

// Default web interface configuration
const DEFAULT_WEB bool = false

// Default web interface port
const DEFAULT_WEB_PORT string = "8888"

// Default slaves for master and servant instances
var DEFAULT_SLAVES []string

func init() {
	name, err := os.Hostname()
	if err != nil {
		DEFAULT_NAME = name
	}
	DEFAULT_TAGS = []string{runtime.GOOS}
	DEFAULT_SLAVES = []string{}
}

// Structure describing the godo instance configuration
//
// Name : instance name
// Type : instance type
// Port : port godo listens on if it's either a slave or a servant
// Tags : job tags accepted by this instance (ONLY APPLICABLE FOR SERVANT OR SLAVE INSTANCES)
// Web : boolean, tells if the instance has its web interface enabled
// WebPort : port for the web interface
// Slaves : list of slaves known to the instance (ONLY APPLICABLE FOR MASTER AND SERVANT INSTANCES)
// SlaveSelectMode : Algorithm used to determine the slave or servant to run the job (ONLY APPLICABLE FOR MASTER AND SERVANT INSTANCES)
type Config struct {
	Name    string
	Type    InstanceType
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

// Apply the configuration of the new config instance to the old one only if
// the new values aren't equal to the default values
func MergeConfig(new Config, old Config) Config {
	return old.MergeConfig(new)
}

// Apply the configuration of the new config instance only if the new values
// aren't equal to the default values
func (config Config) MergeConfig(new Config) Config {
	if new.Name != DEFAULT_NAME {
		config.Name = new.Name
	}
	if new.Type != DEFAULT_TYPE {
		config.Type = new.Type
	}
	if new.Port != DEFAULT_PORT {
		config.Port = new.Port
	}
	if config.Type == Master {
		if reflect.DeepEqual(new.Tags, DEFAULT_TAGS) {
			config.Tags = new.Tags
		}
	}
	if new.Web != DEFAULT_WEB {
		config.Web = new.Web
	}
	if new.WebPort != DEFAULT_PORT {
		config.WebPort = new.WebPort
	}
	if config.Type == Slave || config.Type == Servant {
		if reflect.DeepEqual(new.Slaves, DEFAULT_SLAVES) {
			config.Slaves = new.Slaves
		}
	}
	return config
}

// Constructor for the config struct
func NewConfig() *Config {
	c := new(Config)
	c.Name = DEFAULT_NAME
	c.Type = DEFAULT_TYPE
	c.Port = DEFAULT_PORT
	c.Tags = DEFAULT_TAGS
	c.Web = DEFAULT_WEB
	c.WebPort = DEFAULT_WEB_PORT
	c.Slaves = DEFAULT_SLAVES
	return c
}

// Prints a human readable rundown of the configuration
func (config Config) ToString() string {
	ret := ""
	switch config.Type {
	case Master:
		ret = fmt.Sprintf("[MASTER]")
	case Servant:
		ret = fmt.Sprintf("[SERVANT]")
	case Slave:
		ret = fmt.Sprintf("[SLAVE]")
	}
	return ret
}
