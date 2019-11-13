package fileoperations

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ReadFile returns the contents of a file as []byte
func ReadFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

// WriteFile writes input data to a file
func WriteFile(filename string, data []byte) {
	ioutil.WriteFile(filename, data, 0777)
}

// IsValidFile determines if the filename is valid
func IsValidFile(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
