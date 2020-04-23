package wordchain

import (
	"reflect"
	"testing"
)

func TestWordChain_Chain(t *testing.T) {
	t.Parallel()
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		args      args
		wantWords []string
		wantErr   bool
	}{
		{args{"cat", "dog"}, []string{"cat", "dat", "dot", "dog"}, false},
		{args{"gold", "lead"}, []string{"gold", "goad", "load", "lead"}, false},
		{args{"above", "below"}, []string{"above", "abote", "abate", "alate", "blate", "blats", "boats", "bolts", "boles", "bales", "baler", "balor", "balow", "below"}, false},
		{args{"soup", "rice"}, []string{"soup", "souk", "sock", "rock", "rick", "rice"}, false},
		{args{"ruby", "code"}, []string{"ruby", "rudy", "rude", "rode", "code"}, false},
	}
	// Append also reversed cases e.g. cat-dog becomes dog-cat.
	for _, t := range tests[:len(tests)-1] {
		var wantWords = make([]string, len(t.wantWords))
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

	t.Run("deterministic", func(t *testing.T) {
		t.Parallel()
		w := New()
		// WordChain results are non-deterministic, e.g. dog-cat pair may return [dog cog cot cat] or [dog dot dat cat].
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
		// w.LoadWordsFromFile("/usr/share/dict/words")
		err := w.LoadWordsFromFile("testdata/words_alpha.txt", exclude)
		if err != nil {
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
	})

	t.Run("non-deterministic", func(t *testing.T) {
		t.Parallel()
		// Results may differ each run when we use full dictionary as there are multiple possible outcomes.
		// However there are certain characteristics which we can verify:
		// * length of the chain needs to be the same because algorithm should pick always shortest path
		// * distance between words should always be the same, just one letter difference
		// * first and last word should match initial start and end word
		w := New()
		err := w.LoadWordsFromFile("testdata/words_alpha.txt", map[string]struct{}{})
		if err != nil {
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

				// Length of the chain should be the same for all  because algorithm picks shortest path.
				if len(gotWords) != len(tt.wantWords) {
					t.Errorf("len(Chain()) = %v, want %v", len(gotWords), len(tt.wantWords))
				}

				// Distance between words should always be the same, just one letter difference.
				for i := range gotWords {
					if i == 0 {
						continue
					}

					// An overkill to initialize Word to just use HeuristicScore() func but it's a test 🤷
					w := &Word{
						text: gotWords[i-1],
					}
					expected := 1
					hScore := w.HeuristicScore(gotWords[i])
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
	})
}
