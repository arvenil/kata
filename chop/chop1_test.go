package chop

import (
	"sort"
	"testing"
)

func Test_chop1(t *testing.T) {
	type args struct {
		i int
		v []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Examples taken from http://codekata.com/kata/kata02-karate-chop/ and treated with:
		// %s/assert_equal(\(-\?[0-9]\+\), \+chop(\([0-9]\+\), \[\(.*\)\]))/{"chop1(\2,[\3])", args{\2,[]int{\3}}, \1},/g
		//
		{"chop1(3,[])", args{3, []int{}}, -1},
		{"chop1(3,[1])", args{3, []int{1}}, -1},
		{"chop1(1,[1])", args{1, []int{1}}, 0},
		//
		{"chop1(1,[1, 3, 5])", args{1, []int{1, 3, 5}}, 0},
		{"chop1(3,[1, 3, 5])", args{3, []int{1, 3, 5}}, 1},
		{"chop1(5,[1, 3, 5])", args{5, []int{1, 3, 5}}, 2},
		{"chop1(0,[1, 3, 5])", args{0, []int{1, 3, 5}}, -1},
		{"chop1(2,[1, 3, 5])", args{2, []int{1, 3, 5}}, -1},
		{"chop1(4,[1, 3, 5])", args{4, []int{1, 3, 5}}, -1},
		{"chop1(6,[1, 3, 5])", args{6, []int{1, 3, 5}}, -1},
		//
		{"chop1(1,[1, 3, 5, 7])", args{1, []int{1, 3, 5, 7}}, 0},
		{"chop1(3,[1, 3, 5, 7])", args{3, []int{1, 3, 5, 7}}, 1},
		{"chop1(5,[1, 3, 5, 7])", args{5, []int{1, 3, 5, 7}}, 2},
		{"chop1(7,[1, 3, 5, 7])", args{7, []int{1, 3, 5, 7}}, 3},
		{"chop1(0,[1, 3, 5, 7])", args{0, []int{1, 3, 5, 7}}, -1},
		{"chop1(2,[1, 3, 5, 7])", args{2, []int{1, 3, 5, 7}}, -1},
		{"chop1(4,[1, 3, 5, 7])", args{4, []int{1, 3, 5, 7}}, -1},
		{"chop1(6,[1, 3, 5, 7])", args{6, []int{1, 3, 5, 7}}, -1},
		{"chop1(8,[1, 3, 5, 7])", args{8, []int{1, 3, 5, 7}}, -1},
	}
	for _, tt := range tests {
		// First, sanity check if provided slice is sorted.
		if !sort.IntsAreSorted(tt.args.v) {
			t.Errorf("ints are not sorted: %v", tt.args.v)
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := Classic(tt.args.i, tt.args.v); got != tt.want {
				t.Errorf("chop1() = %v, want %v", got, tt.want)
			}
		})
	}
}
