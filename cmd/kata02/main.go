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
	flag.IntVar(&needle, "needle", 0, "an integer to algorithms for")
	flag.Var(&haystack, "haystack", "comma-separated, sorted, list of integers")
	name := algorithms.Slice[0].Name()
	flag.StringVar(&algorithm, "algorithm", name, fmt.Sprintf("choose from: %s", algorithms.Slice))
}

func main() {
	flag.Parse()

	a, ok := algorithms.Map[algorithm]
	if !ok {
		fmt.Printf("algorithm '%s' doesn't exist, choose from: %s\n", algorithm, algorithms.Slice)
		os.Exit(1)
	}
	fmt.Println(a.Search(needle, haystack))
}
