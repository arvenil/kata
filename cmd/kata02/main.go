package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arvenil/kata/bsearch/algorithms"
)

// Available flags.
var (
	needle    int
	haystack  haystackFlag
	algorithm string
)

func init() {
	flag.IntVar(&needle, "needle", 0, "an integer to search for in haystack e.g. 5")
	flag.Var(&haystack, "haystack", "comma-separated, sorted, list of integers e.g. 1,5,7")
	flag.StringVar(&algorithm, "algorithm", algorithms.Slice[0].String(), fmt.Sprintf("choose from: %s", algorithms.Slice))
}

func main() {
	flag.Parse()

	// At least one flag needs to be passed otherwise print Usage().
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Check if algorithm is valid.
	a, ok := algorithms.Map[algorithm]
	if !ok {
		fmt.Fprintf(os.Stderr, "algorithm '%s' doesn't exist, choose from: %s\n", algorithm, algorithms.Slice)
		os.Exit(1)
	}

	// Search for needle in haystack.
	fmt.Fprintln(os.Stdout, a.Search(needle, haystack))
}
