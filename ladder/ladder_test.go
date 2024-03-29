package ladder_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/arvenil/kata/ladder"
)

type args struct {
	start string
	end   string
}

var tests = []struct {
	args      args
	wantWords []string
	wantErr   bool
}{
	{args{"cat", "dog"}, []string{"cat", "dat", "dot", "dog"}, false},
	{args{"gold", "lead"}, []string{"gold", "goad", "load", "lead"}, false},
	{args{"soup", "rice"}, []string{"soup", "souk", "sock", "rock", "rick", "rice"}, false},
	{args{"ruby", "code"}, []string{"ruby", "rudy", "rude", "rode", "code"}, false},
	{args{"above", "below"}, []string{
		"above",
		"abote",
		"abate",
		"alate",
		"blate",
		"blats",
		"boats",
		"bolts",
		"boles",
		"bales",
		"baler",
		"balor",
		"balow",
		"below",
	}, false},
}

func init() {
	// Push to tests also reversed cases e.g. cat-dog becomes dog-cat.
	for _, t := range tests[:len(tests)-1] {
		wantWords := make([]string, len(t.wantWords))
		for i := range t.wantWords {
			wantWords[i] = t.wantWords[len(t.wantWords)-1-i]
		}

		r := struct {
			args      args
			wantWords []string
			wantErr   bool
		}{
			args{t.args.end, t.args.start},
			wantWords,
			t.wantErr,
		}
		tests = append(tests, r)
	}
}

func TestLadder_Chain_Deterministic(t *testing.T) {
	t.Parallel()

	// Ladder results are non-deterministic, e.g. dog-cat pair may return [dog cog cot cat] or [dog dot dat cat].
	// To visualize common examples in tests it's easier to force only one possibility by excluding some words.
	exclude := map[string]struct{}{
		"cot":   {},
		"cog":   {},
		"dag":   {},
		"rine":  {},
		"sice":  {},
		"rime":  {},
		"rive":  {},
		"sick":  {},
		"roup":  {},
		"rube":  {},
		"ballo": {},
		"abave": {},
		"bolas": {},
		"baloo": {},
		"balon": {},
	}

	// w.Load("/usr/share/dict/words")
	w := ladder.New()
	if err := w.Load("testdata/words_alpha.txt", exclude); err != nil {
		t.Errorf("unable to load word list: %v", err)
	}

	for _, tt := range tests {
		tt := tt
		name := tt.args.start + "-" + tt.args.end
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			gotWords, err := w.Chain(tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chain() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(gotWords, tt.wantWords) {
				t.Errorf("Chain() gotWords = %v, want %v", gotWords, tt.wantWords)
			}
		})
	}
}

func TestLadder_Chain_NonDeterministic(t *testing.T) {
	t.Parallel()
	// Results may differ each run when we use full dictionary as there are multiple possible outcomes.
	// However, there are certain characteristics which we can verify:
	// * length of the ladder needs to be the same because algorithm should pick always the shortest path
	// * distance between words should always be the same, just one letter difference
	// * first and last word should match initial start and end word
	w := ladder.New()
	if err := w.Load("testdata/words_alpha.txt", map[string]struct{}{}); err != nil {
		t.Errorf("unable to load word list: %v", err)
	}

	for _, tt := range tests {
		tt := tt
		name := tt.args.start + "-" + tt.args.end
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			gotWords, err := w.Chain(tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chain() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			// Length of the ladder should be the same for all  because algorithm picks shortest path.
			if len(gotWords) != len(tt.wantWords) {
				t.Errorf("len(Chain()) = %v, want %v", len(gotWords), len(tt.wantWords))
			}

			// Distance between words should always be the same, just one letter difference.
			for i := range gotWords {
				if i == 0 {
					continue
				}

				// An overkill to initialize word to just use Score() algorithms but it's a test 🤷
				w := &ladder.Word{
					Text:           gotWords[i-1],
					Neighbourhoods: nil,
				}
				expected := 1
				hScore := w.Score(gotWords[i])
				if hScore != expected {
					t.Errorf("heuristic score between %v and %v is %v, should be %v", gotWords[i-1], gotWords[i], hScore, expected)
				}
			}

			// First word should match start word.
			if gotWords[0] != tt.args.start {
				t.Errorf("len(Chain()) = %v, want %v", len(gotWords), len(tt.wantWords))
			}

			// Last word should match end word.
			if gotWords[len(gotWords)-1] != tt.args.end {
				t.Errorf("len(Chain()) = %v, want %v", len(gotWords), len(tt.wantWords))
			}
		})
	}
}

func TestLadder_Chain_Errors(t *testing.T) {
	t.Parallel()

	var err error

	w := ladder.New()

	if _, err = w.Chain("cat", "dog"); err == nil {
		t.Error("expected an error: dictionary is not loaded yet")
	}

	if err := w.Load("testdata/words_alpha.txt", map[string]struct{}{}); err != nil {
		t.Errorf("unable to load word list: %v", err)
	}

	if _, err = w.Chain("cat", "mouse"); err == nil {
		t.Error("expected an error: 'start' word has different length than 'end' word")
	}

	if _, err = w.Chain("288D3A55", "288D3A55"); err == nil {
		t.Error("expected an error: word doesn't exist in dictionary")
	}
}

func ExampleLadder_Chain() {
	l := ladder.New()
	if err := l.Load("/usr/share/dict/words", nil); err != nil {
		panic(err)
	}

	words, err := l.Chain("gold", "lead")
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(os.Stdout, words)
	// Output:
	// [gold goad load lead]
}
