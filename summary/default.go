package summary

import (
	"fmt"
	"strings"
)

type Default struct {
	Entries map[string]int
	pattern string
}

func newDefault(pattern string) *Default {
	s := new(Default)
	s.Entries = make(map[string]int)
	s.pattern = pattern
	return s
}

func (d *Default) AddEntry(entry string) {
	entry = d.extractPattern(entry)
	_, ok := d.Entries[entry]
	if ok == false {
		d.Entries[entry] = 1
	} else {
		d.Entries[entry]++
	}
}

func (d *Default) Print() {
	headline := strings.Replace("%s[seperator]%s\n", "[seperator]", seperator, 1)
	fmt.Printf(headline, "Match", "Count")

	output := strings.Replace("%s[seperator]%d\n", "[seperator]", seperator, 1)
	for k, v := range d.Entries {
		fmt.Printf(output, k, v)
	}
}

/**
	extractPattern returns the part of the input string starting with the pattern until next whitespace character
 */
func (d *Default) extractPattern(input string) string {
	start := strings.Index(input, d.pattern)
	if start == -1 {
		return ""
	}

	res := input[start:]
	f := strings.Fields(res)
	if len(f) > 0 {
		return f[0]
	}

	return res
}
