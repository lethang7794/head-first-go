package datafile

import (
	"bufio"
	"os"
)

// GetStrings reads a string from each line of a file.
func GetStrings(fileName string) ([]string, error) {
	var strings []string

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strings = append(strings, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return strings, nil
}
