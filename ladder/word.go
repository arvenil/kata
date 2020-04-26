package ladder

// word as text with its neighbourhoods.
type word struct {
	text           string
	neighbourhoods map[string]map[string]*word
}

// Score is an estimate of the cost to reach word "a".
func (w *word) Score(a string) (score int) {
	for i := range a {
		if a[i] != w.text[i] {
			score++
		}
	}
	return score
}
