package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arvenil/kata/bsearch"
)

// Available flags.
var (
	needle    int
	haystack  haystackFlag
	algorithm string
)

func init() {
	flag.IntVar(&needle, "needle", 0, "an integer to search for")
	flag.Var(&haystack, "haystack", "an ordered, comma separated, list of integers")
	name := bsearch.Algorithms.Sorted()[0].Name()
	flag.StringVar(&algorithm, "algorithm", name, fmt.Sprintf("choose from: %s", bsearch.Algorithms))
}

func main() {
	flag.Parse()

	a, ok := bsearch.Algorithms[algorithm]
	if !ok {
		fmt.Printf("algorithm '%s' doesn't exist, choose from: %s\n", algorithm, bsearch.Algorithms)
		os.Exit(1)
	}
	fmt.Println(a.Search(needle, haystack))
}
