package main

import (
	"github.com/Marneus68/godo/config"
	"github.com/Marneus68/godo/job"
	"testing"
)

func TestNewJob(t *testing.T) {
	c := config.NewConfig()

	j := job.NewJob(&c)
}
