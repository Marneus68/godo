package main

import (
	"github.com/Marneus68/godo/config"
	"testing"
)

func TestNewConfig(t *testing.T) {
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

	if nn.Name != "instance_name" {
		t.Errorf("instance name was not successfully merged")
	}
	if nn.Web != true {
		t.Errorf("instance web service was not successfully merged")

	}
	if nn.Port != "1234" {
		t.Errorf("instance port was not successfully merged")
	}
	if nn.WebPort != "5678" {
		t.Errorf("instance web port was not successfully merged")
	}
}
