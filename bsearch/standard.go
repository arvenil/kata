package bsearch

import (
	"sort"
)

// Standard searches for n (needle) in a sorted slice of ints h (haystack).
// The return value is the index of n or -1 if n is not present in h.
// The slice must be sorted in ascending order.
//
// Standard uses binary algorithms algorithm from Go standard library.
func Standard(n int, h []int) int {
	i := sort.SearchInts(h, n)
	if i < len(h) && h[i] == n {
		return i
	}

	return -1
}
