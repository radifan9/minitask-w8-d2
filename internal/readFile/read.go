package readfile

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func MustReadFile(filePath string) (string, error) {
	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("continue...")
		}
	}()

	file, errOpen := os.Open(filePath)
	if errOpen != nil {
		// Handle error (invalid file)
		return "", errors.New("invalid file path")
	}

	// If the file is not a directory & a valid file path (file can be read)
	scanner := bufio.NewScanner(file)

	var fileContentBuilder strings.Builder
	for scanner.Scan() {
		// Build the string line by line
		line := scanner.Text()
		fileContentBuilder.WriteString(line)
		fileContentBuilder.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("path %s is a directory, not a file", filePath))
	}

	// Close the file
	defer file.Close()

	return fileContentBuilder.String(), nil
}
