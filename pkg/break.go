package pkg

import (
	"os"
	"fmt"
)

const BREAKFILE = "Breakfile"

func getLastModTime(filename string) (string, error) {
	file, err := os.Stat(filename)

	if err != nil {
		return "", err
	}
	return file.ModTime().String(), nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return true
}

func readFile(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
func checkBreakfile(path string) {
	if !fileExists(path) {
		fmt.Fprintf(os.Stderr, "[ERROR]: Breakfile Not Found!\n")
		os.Exit(1)
	}
}

func HandleRoot(directory string, args []string) {
	path := directory + BREAKFILE
	checkBreakfile(path)

	contents, err := readFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR]: Error While Reading File")
		os.Exit(1)
	}

	l := NewLexer{input: contents}
}
