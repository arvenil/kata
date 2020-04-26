package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arvenil/kata/ladder"
)

// Available flags.
var (
	dictionary string
	pairs      pairsFlag
)

func init() {
	flag.StringVar(&dictionary, "d", "/usr/share/dict/words", "path to dictionary")
	flag.Var(&pairs, "p", "two words separated by comma e.g. 'dog,cat'")
}

func main() {
	flag.Parse()
	if len(pairs) == 0 {
		fmt.Println("at least one pair of words is required e.g. '-p dog,cat'")
		os.Exit(1)
	}

	// Create new Ladder and Load the dictionary.
	wc := ladder.New()
	err := wc.Load(dictionary, map[string]struct{}{})
	if err != nil {
		fmt.Printf("can't load word list: %s\n", err)
		os.Exit(1)
	}

	// For each pair of words try to find word-ladder and append it to results.
	results := Results{}
	for _, pair := range pairs {
		words, err := wc.Chain(pair[0], pair[1])
		results.Append(pair[0], pair[1], words, err)
	}

	// Print results.
	fmt.Print(results)
}
