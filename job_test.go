package main

import (
	"fmt"
	"github.com/Marneus68/godo/config"
	"github.com/Marneus68/godo/job"
	"os"
	"strings"
	"testing"
)

func TestNewJob(t *testing.T) {
	fmt.Println("RUNNING JOB TESTS")
	_ = job.NewJob(config.NewConfig())
}

func TestReadJob(t *testing.T) {
	c := config.NewConfig()
	j := job.NewJob(c)
	wd, _ := os.Getwd()
	path := wd + "/testsData/jobs.d/testJob"
	t.Log("TEST DATA PATH: " + path)
	err := j.ReadFromFile(path)
	if err != nil {
		t.Errorf("there was an error reading the job description file")
	}
	if strings.Compare(j.Name, "testData") != 0 {
		t.Error("job name was not successfully read")
	}
}
