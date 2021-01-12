package betterioutil

import (
	"io/ioutil"
)

// ReadFileString ...
func ReadFileString(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	return string(bytes), err
}
