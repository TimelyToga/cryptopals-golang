package set1

import (
	"bufio"
	"os"
	"path/filepath"
)

// ReadLines returns a []string of the lines in a file
func ReadLines(relFilePath string) ([]string, error) {
	absPath, err := filepath.Abs("4.txt")
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
