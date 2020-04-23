package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arvenil/kata/wordchain"
)

// Available flags.
var (
	dictionary string
	pairs      pairsFlag
)

func init() {
	flag.StringVar(&dictionary, "d", "/usr/share/dict/words", "path to word list")
	flag.Var(&pairs, "p", "two words separated by comma e.g. 'dog,cat'")
}

func main() {
	flag.Parse()
	if len(pairs) == 0 {
		fmt.Println("at least one pair of words is required e.g. '-p dog,cat'")
		os.Exit(1)
	}

	// Create new instance of WordChain and load the dictionary.
	wc := wordchain.New()
	err := wc.LoadWordsFromFile(dictionary, map[string]struct{}{})
	if err != nil {
		fmt.Printf("can't load word list: %s\n", err)
		os.Exit(1)
	}

	// For each pair of words try to find word-chain and append it to results.
	results := Results{}
	for _, pair := range pairs {
		words, err := wc.Chain(pair[0], pair[1])
		results.Append(pair[0], pair[1], words, err)
	}

	// Print results.
	fmt.Print(results)
}
