// Based on:
// * https://en.wikipedia.org/wiki/A*_search_algorithm
// *
package wordchain

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func New() *WordChain {
	return &WordChain{
		words: map[string]*Word{},
		masks: map[string]map[string]*Word{},
	}
}

// Word
type Word struct {
	text           string
	neighbourhoods map[string]map[string]*Word
}

// WordChain
type WordChain struct {
	words map[string]*Word
	// [co*d][cord]
	masks map[string]map[string]*Word
}

func (w *WordChain) LoadWordsFromFile(path string, exclude map[string]struct{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if _, ok := exclude[text]; ok {
			continue
		}
		w.Append(text)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (w *WordChain) Append(text string) {
	word := &Word{
		text:           text,
		neighbourhoods: map[string]map[string]*Word{},
	}
	w.words[text] = word

	for i := range text {
		mask := text[:i] + "*" + text[i+1:]
		if w.masks[mask] == nil {
			w.masks[mask] = map[string]*Word{}
		}
		w.masks[mask][text] = word
		word.neighbourhoods[mask] = w.masks[mask]
	}
}

// HeuristicScore is an estimate of the cost to reach word "a" from Word.
func (w *Word) HeuristicScore(a string) (score int) {
	for i := range a {
		if a[i] != w.text[i] {
			score++
		}
	}
	return score
}

func (w *WordChain) Chain(start, end string) (words []string, err error) {
	if len(start) != len(end) {
		return nil, fmt.Errorf("words need to have the same length: %v(%v)-%v(%v)", start, len(start), end, len(end))
	}

	// Lookup starting word in the dictionary.
	word := w.words[start]
	if word == nil {
		return nil, fmt.Errorf("couldn't find the word in dictionary: %v", start)
	}

	// The set of discovered nodes that may need to be (re-)expanded.
	// Initially, only the start node is known.
	// This is usually implemented as a min-heap or priority queue rather than a hash-set.
	openSet := &PriorityQueue{}
	heap.Init(openSet)
	item := &Item{
		word:   word,
		fScore: word.HeuristicScore(end),
	}
	items := map[string]*Item{}
	items[word.text] = item

	heap.Push(openSet, item)

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*Item)
		if current.word.text == end {
			for current != nil {
				words = append([]string{current.word.text}, words...)
				current = current.cameFrom
			}
			return words, nil
		}

		for _, neighbourhood := range current.word.neighbourhoods {
			for _, neighbour := range neighbourhood {

				item, ok := items[neighbour.text]
				if !ok {
					item = &Item{
						gScore: 100000000,
						fScore: 100000000,
					}
					items[neighbour.text] = item
				}
				gScoreTentative := current.gScore + 1
				if gScoreTentative < item.gScore {
					item.word = neighbour
					item.cameFrom = current
					item.gScore = gScoreTentative
					item.fScore = gScoreTentative + neighbour.HeuristicScore(end)
					if !ok {
						heap.Push(openSet, item)
					} else {
						heap.Fix(openSet, item.index)
					}
				}
			}
		}
	}

	return words, nil
}
