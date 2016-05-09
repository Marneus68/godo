package main

import (
	"fmt"
	"github.com/Marneus68/godo/config"
	"github.com/Marneus68/godo/job"
	"testing"
)

func TestNewJob(t *testing.T) {
	fmt.Println("RUNNING JOB TESTS")
	_ = job.NewJob(config.NewConfig())
}

func TestReadJob(t *testing.T) {
	c := config.NewConfig()
	j := job.NewJob(c)
	err := j.ReadFromFile("testData/jobs.d/testJob")
	if err != nil {
		t.Errorf("there was an error reading the job description file")
	}
	if j.Name != "testData" {
		t.Error("job name was not successfully read")
	}
}
