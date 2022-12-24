package main

import (
	"fmt"
	"testing"
)

func TestCheckFileExists(t *testing.T) {
	filename := "./test2"
	fmt.Println(CheckFileExists(filename))
}

func TestAppendExistingFile(t *testing.T) {
	filename := "./testdata"
	err := AppendExistingFile(filename, "\ndef")
	fmt.Println(err)
}

func TestWriteNewFile(t *testing.T) {
	filename := "./testdata"
	err := WriteNewFile(filename, "abc\n123")
	fmt.Println(err)
}

func TestReadCompleteFile(t *testing.T) {
	filename := "./testdata"
	content, err := ReadCompleteFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(content)
}

func TestReadLargeFileLine(t *testing.T) {
	filename := "./testdata"
	err := ReadLargeFileLine(filename)
	fmt.Println(err)
}
