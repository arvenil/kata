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
	flag.Var(&pairs, "p", "two words separated by comma e.g. `dog,cat`")
}

func main() {
	flag.Parse()
	if len(pairs) == 0 {
		fmt.Println("at least one pair of words is required e.g. `-p dog,cat`")
		os.Exit(1)
	}

	wc := wordchain.New()
	err := wc.LoadWordsFromFile(dictionary, map[string]struct{}{})
	if err != nil {
		fmt.Printf("can't load word list: %s\n", err)
		os.Exit(1)
	}
	for _, pair := range pairs {
		var result string
		words, err := wc.Chain(pair[0], pair[1])
		if err != nil {
			result = err.Error()
		} else {
			result = fmt.Sprintf("%v", words)
		}
		fmt.Printf("%v-%v: %v\n", pair[0], pair[1], result)
	}
}
