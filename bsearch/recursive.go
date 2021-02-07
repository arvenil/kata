package bsearch

// Recursive searches for n (needle) in a sorted slice of ints h (haystack).
// The return value is the index of n or -1 if n is not present in h.
// The slice must be sorted in ascending order.
//
// Recursive uses recursion to implement binary algorithms.
func Recursive(n int, h []int) int {
	return recursive(n, h, 0, len(h)-1)
}

func recursive(n int, h []int, l, r int) int {
	if l > r {
		return -1
	}

	m := (l + r) / 2 //nolint:gomnd
	if h[m] == n {
		return m
	}

	if n < h[m] {
		return recursive(n, h, l, m-1)
	}

	return recursive(n, h, m+1, r)
}
