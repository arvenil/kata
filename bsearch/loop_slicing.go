package bsearch

// LoopSlicing searches for n (needle) in a sorted slice of ints h (haystack).
// The return value is the index of n or -1 if n is not present in h.
// The slice must be sorted in ascending order.
//
// LoopSlicing uses 'for' loop to implement binary algorithms.
//
// LoopSlicing slices haystack to new boundaries every iteration of 'for' loop.
func LoopSlicing(n int, h []int) (i int) {
	for len(h) > 0 {
		m := (len(h) - 1) / 2 //nolint:gomnd

		switch {
		case h[m] < n:
			h = h[m+1:]
			i += m + 1
		case h[m] > n:
			h = h[:m]
		default:
			return m + i
		}
	}

	// Return -1 when not found.
	return -1
}
