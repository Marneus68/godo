// Package keyValueParser implements a simple key=value parser
package keyValueParser

import (
	"os"
)

func readFile(filename string) (map[string]string, error) {
	ret = map[string]string{}
	f, err := os.Open(filename)
	if err != nil {
		return ret, err
	}
	defer f.Close()

	// TODO

	return ret, err
}
