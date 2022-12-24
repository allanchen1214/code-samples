package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Check if file exists
func CheckFileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// Read a complete file
func ReadCompleteFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
	//	lines = strings.Split(string(content), "\n")
}

// Read a large file line-by-line
func ReadLargeFileLine(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Use bufio scanner, the default Scan method is by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line) // line content
	}
	return nil
}

// Write a new file
func WriteNewFile(filename, content string) error {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

// Append a existing file
func AppendExistingFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println(filepath.Join("a", "b", "file.ext"))
}
