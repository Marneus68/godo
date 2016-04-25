// Package containing the definition of the configuration
package config

import (
	"fmt"
	"github.com/Marneus68/godo/utils"
	"os"
	"path/filepath"
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

// Enumeration describing the way slaves or servants are selected to submit a
// job
type SlaveSelectMode int

const (
	RR        SlaveSelectMode = iota // Round robin between all slaves
	RR_W_SELF                        // Round robin between all slaves and the current instance
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

/*
Structure describing the godo instance configuration

	Name : instance name
	Type : instance type
	Port : port godo listens on if it's either a slave or a servant
	Tags : job tags accepted by this instance (ONLY APPLICABLE FOR SERVANT OR SLAVE INSTANCES)
	Web : boolean, tells if the instance has its web interface enabled
	WebPort : port for the web interface
	Slaves : list of slaves known to the instance (ONLY APPLICABLE FOR MASTER AND SERVANT INSTANCES)
	SlaveSelectMode : Algorithm used to determine the slave or servant to run the job (ONLY APPLICABLE FOR MASTER AND SERVANT INSTANCES)
*/
type Config struct {
	Name    string
	Type    InstanceType
	Port    string
	Tags    []string
	Web     bool
	WebPort string
	Slaves  []string
}

// Default configuration file
const CONF_FILE_NAME string = "godo.conf"

// Default name of the jobs directory
const JOBS_DIR_NAME string = "jobs.d"

// Default location of the configuration file and job subdirectories
const DEFAULT_CONF_LOCATION string = "~/godo/"

var locations = []string{
	"/godo/",
	"/etc/godo/",
	"/usr/local/share/godo/",
	"~/.config/godo/",
	"~/godo/",
}

var config string = ""
var configFile string = ""
var jobsDirectory string = ""

// Looks for the configuration directory from multiple possible
// locations, returns the full absolute path to the first
// configuration file found
func ConfigDirectory() string {
	if strings.Compare(config, "") != 0 {
		return config
	}
	for _, v := range locations {
		v = utils.SubstituteHomeDir(v)
		fmt.Println(v)
		_, err := os.Stat(v)
		if err == nil {
			fmt.Print("FOUND")
			config = utils.SubstituteHomeDir(v)
			break
		}
	}
	if strings.Compare(config, "") == 0 {
		//log.Fatal("Could not find configuration file.")
		fmt.Println("Could not find a configuration directory...")
		fmt.Println("Defaulting to \"" + DEFAULT_CONF_LOCATION + "\"...")
		p := utils.SubstituteHomeDir(DEFAULT_CONF_LOCATION)
		_ = os.MkdirAll(p, 0777)
		config = p
	}
	return config
}

// Returns the absolute path to the local godo configuration file
func ConfigFile() string {
	if strings.Compare(configFile, "") != 0 {
		return configFile
	}
	configFile = filepath.Join(ConfigDirectory(), CONF_FILE_NAME)
	return configFile
}

// Returns the absolute path to the local jobs directory
func JobsDirectory() string {
	if strings.Compare(jobsDirectory, "") != 0 {
		return jobsDirectory
	}
	jobsDirectory = filepath.Join(ConfigDirectory(), JOBS_DIR_NAME)
	return jobsDirectory
}

// Read configuration from file
func (config Config) ReadFromFile(path string) {

}

// Save configuration to file
func (config Config) SaveToFile(path string) {

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
