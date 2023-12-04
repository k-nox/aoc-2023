package util

import (
	"bufio"
	"os"
)

type FileScanner struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewFileScanner(path string) *FileScanner {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return &FileScanner{
		file:    file,
		scanner: bufio.NewScanner(file),
	}
}

func (f *FileScanner) Close() {
	err := f.file.Close()
	if err != nil {
		panic(err)
	}
}

func (f *FileScanner) Scan() bool {
	return f.scanner.Scan()
}

func (f *FileScanner) Text() string {
	return f.scanner.Text()
}
