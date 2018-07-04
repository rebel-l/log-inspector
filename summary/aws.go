package summary

import (
	"fmt"
	"strings"
	"strconv"
)

type Aws struct {
	Entries map[string]*AwsEntry
}

func newAws() *Aws {
	a := new(Aws)
	a.Entries = make(map[string]*AwsEntry)
	return a
}

func (a *Aws) AddEntry(entry string) {
	f := strings.Fields(entry)
	if len(f) != 26 {
		return
	}

	uri := f[7]
	code, _ := strconv.Atoi(f[8])
	_, ok := a.Entries[uri]
	if ok == false {
		a.Entries[uri] = newAwsEntry(uri, code, f[10])
	} else {
		a.Entries[uri].Count++
		a.Entries[uri].AddStatusCode(code)
		a.Entries[uri].AddUserAgent(f[10])
	}
}

func (a *Aws) Print() {
	headline := strings.Replace("%s[seperator]%s[seperator]%s[seperator]%s\n", "[seperator]", seperator, 3)
	fmt.Printf(headline, "URI", "Count", "StatusCodes", "UserAgents")

	output := strings.Replace("%s[seperator]%d[seperator]%s[seperator]%s\n", "[seperator]", seperator, 3)

	for _, v := range a.Entries {
		fmt.Printf(output, v.Uri, v.Count, v.GetStatusCodes(), v.GetUserAgents())
	}
}
