package bsearch

// RecursiveSlicing searches for n (needle) in a sorted slice of ints h (haystack).
// The return value is the index of n or -1 if n is not present in h.
// The slice must be sorted in ascending order.
//
// RecursiveSlicing uses recursion to implement binary algorithms.
//
// RecursiveSlicing calls itself with haystack sliced to new boundaries.
func RecursiveSlicing(n int, h []int) int {
	m := len(h) / 2 //nolint:gomnd

	switch {
	case len(h) == 0:
		return -1 // not found
	case h[m] > n:
		return RecursiveSlicing(n, h[:m])
	case h[m] < n:
		i := RecursiveSlicing(n, h[m+1:])
		if i == -1 {
			return -1
		}

		return i + m + 1
	}

	return m
}
