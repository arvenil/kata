package bsearch

import (
	"math/rand"
	"sort"
	"testing"
)

type args struct {
	i int
	v []int
}
type test struct {
	name string
	args args
	want int
}

// Examples taken from http://codekata.com/kata/kata02-karate-chop/ and treated with:
// %s/assert_equal(\(-\?[0-9]\+\), \+bsearch(\([0-9]\+\), \[\(.*\)\]))/{"bsearch(\2,[\3])", args{\2,[]int{\3}}, \1},/g
var tests = []test{
	{"bsearch(3,[])", args{3, []int{}}, -1},
	{"bsearch(3,[1])", args{3, []int{1}}, -1},
	{"bsearch(1,[1])", args{1, []int{1}}, 0},
	//
	{"bsearch(1,[1, 3, 5])", args{1, []int{1, 3, 5}}, 0},
	{"bsearch(3,[1, 3, 5])", args{3, []int{1, 3, 5}}, 1},
	{"bsearch(5,[1, 3, 5])", args{5, []int{1, 3, 5}}, 2},
	{"bsearch(0,[1, 3, 5])", args{0, []int{1, 3, 5}}, -1},
	{"bsearch(2,[1, 3, 5])", args{2, []int{1, 3, 5}}, -1},
	{"bsearch(4,[1, 3, 5])", args{4, []int{1, 3, 5}}, -1},
	{"bsearch(6,[1, 3, 5])", args{6, []int{1, 3, 5}}, -1},
	//
	{"bsearch(1,[1, 3, 5, 7])", args{1, []int{1, 3, 5, 7}}, 0},
	{"bsearch(3,[1, 3, 5, 7])", args{3, []int{1, 3, 5, 7}}, 1},
	{"bsearch(5,[1, 3, 5, 7])", args{5, []int{1, 3, 5, 7}}, 2},
	{"bsearch(7,[1, 3, 5, 7])", args{7, []int{1, 3, 5, 7}}, 3},
	{"bsearch(0,[1, 3, 5, 7])", args{0, []int{1, 3, 5, 7}}, -1},
	{"bsearch(2,[1, 3, 5, 7])", args{2, []int{1, 3, 5, 7}}, -1},
	{"bsearch(4,[1, 3, 5, 7])", args{4, []int{1, 3, 5, 7}}, -1},
	{"bsearch(6,[1, 3, 5, 7])", args{6, []int{1, 3, 5, 7}}, -1},
	{"bsearch(8,[1, 3, 5, 7])", args{8, []int{1, 3, 5, 7}}, -1},
}

func TestSearch(t *testing.T) {
	for _, tt := range tests {
		// First, sanity check if provided slice is sorted.
		if !sort.IntsAreSorted(tt.args.v) {
			t.Errorf("ints are not sorted: %v", tt.args.v)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := Loop(tt.args.i, tt.args.v); got != tt.want {
				t.Errorf("chop1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSearch(b *testing.B) {
	// Generate large set for benchmark.
	rand.Seed(42)
	var largeSet []int
	n := 0
	for i := 1; i < 1000000; i++ {
		n += rand.Intn(10)
		largeSet = append(largeSet, n)
	}
	tests := append(tests, test{"bsearch(4686,[large-set])", args{4686, largeSet}, 998})

	for _, a := range Algorithms.Sorted() {
		b.Run(a.Name(), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for _, tt := range tests {
					if got := a.Search(tt.args.i, tt.args.v); got != tt.want {
						b.Errorf("Search(%v) = %v, want %v", tt.args.i, got, tt.want)
					}
				}
			}
		})
	}
}
