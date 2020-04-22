package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arvenil/kata/chop"
)

// Available flags.
var (
	needle    int
	haystack  haystackFlag
	algorithm string
)

var algorithms map[string]chop.Chopper

func init() {
	flag.IntVar(&needle, "needle", 0, "an integer to search for")
	flag.Var(&haystack, "haystack", "an ordered, comma separated, list of integers")
	flag.StringVar(&algorithm, "algorithm", "classic", fmt.Sprintf("algorithm to use: %v", algorithms))

	algorithms = map[string]chop.Chopper{
		"classic":  chop.ChopFunc(chop.Classic),
		"standard": chop.ChopFunc(chop.Standard),
	}
}

func main() {
	flag.Parse()

	if _, ok := algorithms[algorithm]; !ok {
		fmt.Printf("algorithm '%v' doesn't exist, pick one from: %v\n", algorithm, algorithms)
		os.Exit(1)
	}

	a := algorithms[algorithm]
	fmt.Println(a.Chop(needle, haystack))
}
