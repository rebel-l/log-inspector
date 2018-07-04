package summary

import "fmt"

type Aws struct {

}

func newAws() *Aws {
	a := new(Aws)
	return a
}

func (a *Aws) AddEntry(entry string) {

}

func (a *Aws) Print() {
	fmt.Println("AWS")
}

// Summary: uri, num, statusCodes, userAgent