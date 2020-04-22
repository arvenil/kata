package main

import (
	"fmt"
	"strings"
)

type pairsFlag [][]string

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (p *pairsFlag) String() string {
	return fmt.Sprint(*p)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (p *pairsFlag) Set(value string) error {
	words := strings.Split(value, ",")
	if len(words) != 2 {
		return fmt.Errorf("flag requires exactly two words separated by comma e.g. `dog,cat` but got: %v", value)
	}
	*p = append(*p, words)
	return nil
}
