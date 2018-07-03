package main

import (
	"os"
	"compress/gzip"
	"bufio"
)

type File struct {
	Name string
}

func NewFile(name string) *File {
	f := new(File)
	f.Name = name
	return f
}

func (f *File) Process(pattern string) {
	ff, err := os.Open(f.Name)
	defer ff.Close()

	b := bufio.NewReader(ff)

	r, err := gzip.NewReader(b)
	if err != nil {
		return
	}

	match(pattern, bufio.NewReader(r))
}
