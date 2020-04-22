package chop

import (
	"sort"
)

type Chopper interface {
	Chop(i int, v []int) int
}

type ChopFunc func(i int, v []int) int

func (f ChopFunc) Chop(i int, v []int) int {
	return f(i, v)
}

func Classic(i int, v []int) int {
	l := 0
	r := len(v) - 1

	for l <= r {
		m := (l + r) / 2
		if v[m] < i {
			l = m + 1
		} else if v[m] > i {
			r = m - 1
		} else {
			return m
		}
	}

	// Return -1 when not found.
	return -1
}

func Standard(i int, v []int) int {
	// This is also a definition of `sort.SearchInts(a []int, x int) int`.
	return sort.Search(len(v), func(j int) bool { return v[i] >= j })
}
