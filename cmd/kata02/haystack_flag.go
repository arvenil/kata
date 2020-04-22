package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type haystackFlag []int

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (h *haystackFlag) String() string {
	return fmt.Sprint(*h)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (h *haystackFlag) Set(value string) error {
	// note: this could be improved by rewriting strings.Split to return []int directly instead of []string
	words := strings.Split(value, ",")
	for _, word := range words {
		i, err := strconv.Atoi(word)
		if err != nil {
			return fmt.Errorf("%v is not an integer", word)
		}
		*h = append(*h, i)
	}

	if !sort.IntsAreSorted(*h) {
		return fmt.Errorf("integers are not sorted: %v", *h)
	}

	return nil
}
