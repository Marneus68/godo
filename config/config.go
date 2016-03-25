package config

//import "os"
//import "path"
//import "strings"
import "runtime"

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
}

// Look for the configuration file from multiple possible locations
// returns the full path to the first configuration file found
func FindConfigFile() string {
    return ""
}

// Read configuration from a config file 
func (config Config) FromConfigFile() {

}

// Read configuration from a command line parameters
func (config Config) FromArgs() {

}

// Constructor for the config struct
func newConfig() *Config {
    c := new(Config)
    c.Type = Master
    c.Os = runtime.GOOS
    c.Port = "8008"
    c.Tags = make([]string, 0)
    c.Web = false
    c.WebPort = "8888"
    return c
}

