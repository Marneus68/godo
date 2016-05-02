package main

import (
	"fmt"
	"github.com/Marneus68/godo/config"
	"github.com/Marneus68/godo/job"
	"testing"
)

func TestNewJob(t *testing.T) {
	fmt.Println("RUNNING JOB TESTS")
	j := job.NewJob(config.NewConfig())
	j.Run()
}
