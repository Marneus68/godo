package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCanary(t *testing.T) {
	fmt.Println("RUNNING CANARY TESTS")
	if strings.Compare(os.Getenv("GOPATH"), "") == 0 {
		t.Error("GOPATH not set!")
	}

	_, err := os.Stat(os.Getenv("GOPATH"))
	if os.IsNotExist(err) {
		t.Error("GOPATH doesn't exist!")
	}
	if os.IsPermission(err) {
		t.Error("GOPATH can't be written or read!")
	}

	if !strings.Contains(os.Getenv("PATH"), os.Getenv("GOPATH")) {
		t.Skip("PATH doesn't contain GOPATH's bin directory, you might run into trouble running \"go install\"...")
	}
}
