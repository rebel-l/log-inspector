package main

import "fmt"

type Summary struct {
	Entries map[string]int
}

func NewSummary() *Summary {
	s := new(Summary)
	s.Entries = make(map[string]int)
	return s
}

func (s *Summary) AddEntry(entry string) {
	_, ok := s.Entries[entry]
	if ok == false {
		s.Entries[entry] = 1
	} else {
		s.Entries[entry]++
	}
}

func (s *Summary) Print() {
	for k, v := range s.Entries {
		fmt.Printf("%s | %d\n", k, v)
	}
}
