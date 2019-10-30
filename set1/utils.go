package set1

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadLines returns a []string of the lines in a file
func ReadLines(relFilePath string) ([]string, error) {
	absPath, err := filepath.Abs(relFilePath)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Collect all the lines into a slice and return it
	var output []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// This shouldn't be slow
		// Source: https://www.reddit.com/r/golang/comments/4rinrk/how_do_you_create_an_array_of_variable_length/d51dr0q/
		output = append(output, scanner.Text())
	}

	return output, nil
}

// ReadFileBytes reads a file and returns it as a byte array
func ReadFileBytes(relFilePath string) ([]byte, error) {
	absPath, err := filepath.Abs(relFilePath)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file in
	fileBytes, err := ioutil.ReadAll(file)

	// Base64 decode input
	decodedBytes, err := base64.StdEncoding.DecodeString(string(fileBytes))
	if err != nil {
		return nil, err
	}

	return decodedBytes, nil
}

// MinOf is an int min helper function
// Source: https://stackoverflow.com/a/53709517
func MinOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}
