package utils

import (
	"log"
	"os/user"
	"strings"
)

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
		path = strings.Replace(path, "~/", homeDir, 1)
	}
	return path
}
