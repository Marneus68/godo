// Package utils implements various global utility functions
package utils

import (
	"bufio"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

//
func IsValidPortString(port string) bool {
	p, err := strconv.Atoi(port)
	if err != nil {
		return false
	}
	if p > 65535 {
		return false
	}
	return true
}

// Substitures the tilde (~) character for the home directory of the
// current user (but only if its the first character of the string)
func SubstituteHomeDir(path string) string {
	u, err := user.Current()
	if err != nil {
		log.Fatal("Unable to rerieve the current user's information.")
	}
	homeDir := u.HomeDir
	if homeDir == "" {
		log.Fatal("Unable to find the home directory of the current user.")
	}
	if path[:2] == "~/" {
		//path = strings.Replace(path, "~/", "", 1)
		path = filepath.Join(homeDir, path[2:])
	}
	return filepath.Clean(path)
}

func ReadKeyValueFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ParseKeyValueFile(filename string) (ret map[string]string, err error) {
	lines, err := ReadKeyValueFile(filename)
	if err != nil {
		return ret, err
		//log.Fatalf("readLines: %s", err)
	}
	for _, line := range lines {
		split := strings.Split(line, "=")
		if len(split) == 2 {
			ret[strings.TrimSpace(split[0])] = strings.TrimSpace(split[1])
		}
	}
	return ret, err
}
