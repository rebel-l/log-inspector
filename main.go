package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/rebel-l/log-inspector/summary"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const modeStream = "stream"
const modeFile = "file"

var pattern *string
var directory *string
var aws *bool
var s summary.Summary

func main() {
	if initFlags() == false {
		return
	}

	if *aws {
		s = summary.New(*pattern, summary.StyleAws)
	} else {
		s = summary.New(*pattern)
	}

	info, _ := os.Stdin.Stat()
	switch detectMode(info) {
	case modeStream:
		processStream(info)
		break
	case modeFile:
		processFiles(*directory)
		break
	default:
		return
	}
	s.Print()
}

func match(pattern string, reader *bufio.Reader) {
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}

		if strings.Contains(input, pattern) {
			s.AddEntry(input)
		}
	}
}

func initFlags() bool {
	pattern = flag.String("pattern", "", "Pattern definition to look for")
	directory = flag.String("dir", "", "Path of the files to inspect")
	aws = flag.Bool("aws", false, "Flag indicates to parse AWS CloudFront logs")
	flag.Parse()

	if *pattern == "" {
		fmt.Println("Pattern argument is missing.")
		fmt.Println("Usage:")
		flag.PrintDefaults()
		return false
	}

	return true
}

func detectMode(fi os.FileInfo) string {
	if *directory != "" {
		_, err := os.Stat(*directory)
		if err == nil {
			return modeFile
		}
	}

	if (fi.Mode() & os.ModeCharDevice) != os.ModeCharDevice {
		return modeStream
	}

	fmt.Println("The command is intended to work with pipes or files.")
	fmt.Println("Usage:")
	fmt.Println("  cat yourfile.txt | log-inspector -pattern=<your_pattern>")
	fmt.Println("  log-inspector -pattern=<your_pattern> -dir=<directory-of-your-files>")

	return ""
}



func processStream(info os.FileInfo)  {
	if info.Size() > 0 {
		reader := bufio.NewReader(os.Stdin)
		match(*pattern, reader)
	}
}

func processFiles(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("Could not read directory: %s\n", err.Error())
		return
	}

	for _, f := range files {
		fileMatcher := NewFile(dir + "/" + f.Name())
		fileMatcher.Process(*pattern)
	}
}
