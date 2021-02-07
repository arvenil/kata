package algorithms_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/arvenil/kata/bsearch/algorithms"
)

type args struct {
	i int
	v []int
}

type test struct {
	args args
	want int
}

// Examples taken from http://codekata.com/kata/kata02-karate-chop/ and treated with:
//   %s/assert_equal(\(-\?[0-9]\+\), \+chop(\([0-9]\+\), \[\(.*\)\]))/{args{\2,[]int{\3}}, \1},/g
var tests = []test{
	{args{3, []int{}}, -1},
	{args{3, []int{1}}, -1},
	{args{1, []int{1}}, 0},
	{args{1, []int{1, 3, 5}}, 0},
	{args{3, []int{1, 3, 5}}, 1},
	{args{5, []int{1, 3, 5}}, 2},
	{args{0, []int{1, 3, 5}}, -1},
	{args{2, []int{1, 3, 5}}, -1},
	{args{4, []int{1, 3, 5}}, -1},
	{args{6, []int{1, 3, 5}}, -1},
	{args{1, []int{1, 3, 5, 7}}, 0},
	{args{3, []int{1, 3, 5, 7}}, 1},
	{args{5, []int{1, 3, 5, 7}}, 2},
	{args{7, []int{1, 3, 5, 7}}, 3},
	{args{0, []int{1, 3, 5, 7}}, -1},
	{args{2, []int{1, 3, 5, 7}}, -1},
	{args{4, []int{1, 3, 5, 7}}, -1},
	{args{6, []int{1, 3, 5, 7}}, -1},
	{args{8, []int{1, 3, 5, 7}}, -1},
}

func TestSearch(t *testing.T) {
	t.Parallel()

	for _, algorithm := range algorithms.Slice {
		algorithm := algorithm

		for _, tt := range tests {
			tt := tt
			// First, sanity check if provided slice is sorted.
			if !sort.IntsAreSorted(tt.args.v) {
				t.Errorf("ints are not sorted: %v", tt.args.v)
			}

			name := fmt.Sprintf("%s(%v,%v)", algorithm, tt.args.i, tt.args.v)
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				if got := algorithm.Search(tt.args.i, tt.args.v); got != tt.want {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			})
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	// Generate large slice for benchmark.
	rand.Seed(42)

	var (
		largeSet []int
		n        = 0
	)

	for i := 1; i < 1000000; i++ {
		n += rand.Intn(10) //nolint:gosec
		largeSet = append(largeSet, n)
	}

	tests := append(tests, test{args{4686, largeSet}, 998})

	for _, algorithm := range algorithms.Slice {
		algorithm := algorithm
		b.Run(algorithm.String(), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for _, tt := range tests {
					if got := algorithm.Search(tt.args.i, tt.args.v); got != tt.want {
						b.Errorf("%s(%v) = %v, want %v", algorithm, tt.args.i, got, tt.want)
					}
				}
			}
		})
	}
}
