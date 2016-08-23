// Package job implements the job definition and its association functions
package job

import (
	"github.com/Marneus68/godo/config"
	"github.com/Marneus68/utils"
	"path/filepath"
	"strings"
	//"os"
)

// Structure describing a godo job
//
//    Name: name of the job
//    Tags: job tags expected by this job to be run by a slave or servant instance
//    Command: command executed by the job
//    Config: pointer to the confguration of the current instace if it's a slave or servant
type Job struct {
	Name    string
	Tags    []string
	Command string
	Config  *config.Config
}

// Creates a new empty job
func NewJob(config *config.Config) *Job {
	j := new(Job)
	j.Name = ""
	j.Tags = make([]string, 0)
	j.Command = ""
	j.Config = config
	return j
}

// Creates a new job from a file
func NewJobFromFile(path string, config *config.Config) *Job {
	j := new(Job)
	j.Name = ""
	j.Tags = make([]string, 0)
	j.Command = ""
	j.Config = config
	return j
}

// Read job from file
func ReadFromFile(path string) (j Job, err error) {
	kv, err := utils.ParseKeyValueFile(path)
	if err != nil {
		return j, err
	}
	j.Name = filepath.Base(path)
	if t, ok := kv["tags"]; ok {
		tags := strings.Split(t, ",")
		for i, s := range tags {
			tags[i] = strings.TrimSpace(s)
		}
		j.Tags = tags
	}
	if n, ok := kv["command"]; ok {
		j.Command = strings.TrimSpace(n)
	}
	return j, nil
}

// Save job to file
func (job Job) SaveToFile(path string) {

}

func (job Job) Run() {
	go func() {
		////p, err := os.StarProcess(
	}()
}

func Jobs(path string) ([]Job, error) {
	return nil, nil
}
