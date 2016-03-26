package config_test

import (
	//"fmt"
	//"strconv"
	"godo/config"
	"testing"
)

func TestCoverage(t *testing.T) {
	/*
		t.Log() // displays something
		t.Fail() // marks as failed but continues
		t.FailNow() // marks as failed and stopes
		t.Error() // Log + Fail
		t.Fatal() // Log + FailNow
	*/
	c := config.NewConfig()

	if c.Name == "" {
		t.Log("Possible error while getting the hostname, you may be on an unsupported platform.")
	}
}

func TestConfig(t *testing.T) {

}
