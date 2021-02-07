package ladder

// Word as Text with its Neighbourhoods.
type Word struct {
	Text           string
	Neighbourhoods map[string]map[string]*Word
}

// Score is an estimate of the cost to reach word "a".
func (w *Word) Score(a string) (score int) {
	for i := range a {
		if a[i] != w.Text[i] {
			score++
		}
	}

	return score
}
