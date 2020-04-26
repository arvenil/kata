package bsearch

// Loop searches for n (needle) in a sorted slice of ints h (haystack).
// The return value is the index of n or -1 if n is not present in h.
// The slice must be sorted in ascending order.
//
// Loop uses 'for' loop to implement binary algorithms.
func Loop(n int, h []int) int {
	l := 0
	r := len(h) - 1
	return loop(n, h, l, r)
}

func loop(n int, h []int, l, r int) int {
	for l <= r {
		m := (l + r) / 2
		if h[m] < n {
			l = m + 1
		} else if h[m] > n {
			r = m - 1
		} else {
			return m
		}
	}

	// Return -1 when not found.
	return -1
}
