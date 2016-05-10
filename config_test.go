package main

import (
	"fmt"
	"github.com/Marneus68/godo/config"
	"os"
	"strings"
	"testing"
)

func TestNewConfig(t *testing.T) {
	fmt.Println("RUNNING CONFIG TESTS")
	c := config.NewConfig()

	if c.Name != config.DEFAULT_NAME {
		t.Errorf("instance name doesn't look right")
	}
	if c.Type != config.DEFAULT_TYPE {
		t.Errorf("instance type doesn't look right")
	}
	if c.Port != config.DEFAULT_PORT {
		t.Errorf("instance port doesn't look right")
	}
	if c.Web != config.DEFAULT_WEB {
		t.Errorf("instance web interface status doesn't look right")
	}
	if c.WebPort != config.DEFAULT_WEB_PORT {
		t.Errorf("instance web port doesn't look right")
	}
}

func TestMergeConfig(t *testing.T) {
	c := config.NewConfig()
	n := config.NewConfig()

	n.Name = "instance_name"
	n.Web = true
	n.Port = "1234"
	n.WebPort = "5678"

	nn := c.MergeConfig(*n)

	if strings.Compare(nn.Name, "instance_name") != 0 {
		t.Errorf("instance name was not successfully merged")
	}
	if nn.Web != true {
		t.Errorf("instance web service was not successfully merged")

	}
	if strings.Compare(nn.Port, "1234") != 0 {
		t.Errorf("instance port was not successfully merged")
	}
	if strings.Compare(nn.WebPort, "5678") != 0 {
		t.Errorf("instance web port was not successfully merged")
	}
}

func TestReadConfig(t *testing.T) {
	wd, _ := os.Getwd()
	path := wd + "/testsData/godo.config"
	t.Log("TEST DATA PATH: " + path)
	c, err := config.ReadFromFile(path)
	if err != nil {
		t.Errorf("there was an error opening the test data")
	}

	if strings.Compare(c.Name, "test") != 0 {
		t.Errorf("instance name was not successfully read")
	}
	if c.Type != config.Master {
		t.Errorf("instance type was not successfully read")
	}
	if strings.Compare(c.Port, "1234") != 0 {
		t.Errorf("instance port was not successfully read")
	}
	if c.Web != true {
		t.Errorf("instance web status was not successfully read")
	}
	if strings.Compare(c.WebPort, "5678") != 0 {
		t.Errorf("instance web port was not successfully read")
	}
}
