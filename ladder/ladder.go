package ladder

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// New instance of *Ladder.
func New() *Ladder {
	return &Ladder{}
}

// Ladder is a word-ladder puzzle solver.
type Ladder struct {
	words map[string]*Word
	// [co*d][cord]
	masks map[string]map[string]*Word
}

// Chain of words from start to end, in which two adjacent words
// (that is, words in successive steps) differ by one letter.
func (l *Ladder) Chain(start, end string) (words []string, err error) {
	if len(start) != len(end) {
		return nil, fmt.Errorf("words '%v'(%v) and '%v'(%v) have different length", start, len(start), end, len(end))
	}

	// Be sure dictionary is loaded.
	if !l.loaded() {
		return nil, fmt.Errorf("no dictionary, please Load dictionary first")
	}

	// Lookup starting word in the dictionary.
	word := l.words[start]
	if word == nil {
		return nil, fmt.Errorf("could not find '%v' in dictionary", start)
	}

	// The set of discovered nodes that may need to be (re-)expanded.
	// Initially, only the start node is known.
	openSet := &priorityQueue{}
	heap.Init(openSet)

	item := &node{
		word:   word,
		fScore: word.Score(end),
	}
	items := map[string]*node{}
	items[word.Text] = item

	heap.Push(openSet, item)

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*node)
		if current.word.Text == end {
			for current != nil {
				words = append([]string{current.word.Text}, words...)
				current = current.prev
			}

			return words, nil
		}

		for _, neighbourhood := range current.word.Neighbourhoods {
			for _, neighbour := range neighbourhood {
				item, ok := items[neighbour.Text]
				if !ok {
					item = &node{
						gScore: 100000000,
						fScore: 100000000,
					}
					items[neighbour.Text] = item
				}

				gScoreTentative := current.gScore + 1
				if gScoreTentative < item.gScore {
					item.word = neighbour
					item.prev = current
					item.gScore = gScoreTentative
					item.fScore = gScoreTentative + neighbour.Score(end)

					if !ok {
						heap.Push(openSet, item)
					} else {
						heap.Fix(openSet, item.index)
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("could not find word ladder")
}

// Load dictionary from path but without words on exclude list.
func (l *Ladder) Load(path string, exclude map[string]struct{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	l.words = map[string]*Word{}
	l.masks = map[string]map[string]*Word{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if _, ok := exclude[text]; ok {
			continue
		}

		l.push(text)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// push word w into dictionary.
func (l *Ladder) push(w string) {
	l.words[w] = &Word{
		Text:           w,
		Neighbourhoods: map[string]map[string]*Word{},
	}

	for i := range w {
		mask := w[:i] + "*" + w[i+1:]
		if l.masks[mask] == nil {
			l.masks[mask] = map[string]*Word{}
		}

		l.masks[mask][w] = l.words[w]
		l.words[w].Neighbourhoods[mask] = l.masks[mask]
	}
}

// loaded returns true if dictionary was already loaded.
func (l *Ladder) loaded() bool {
	switch {
	case l.words == nil, l.masks == nil:
		return false
	}

	return true
}
