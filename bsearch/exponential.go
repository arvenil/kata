package bsearch

// Exponential searches for n (needle) in a sorted slice of ints h (haystack).
// The return value is the index of n or -1 if n is not present in h.
// The slice must be sorted in ascending order.
//
// Exponential algorithms extends binary algorithms to unbounded lists.
// It starts by finding the first element with an index that is both a power of two and greater than the target value.
// Afterwards, it sets that index as the upper bound, and switches to binary algorithms.
// Exponential algorithms works on bounded lists, but becomes an improvement over binary algorithms
// only if the target value lies near the beginning of the array.
//
// https://en.wikipedia.org/wiki/Exponential_search
func Exponential(n int, h []int) int {
	if len(h) == 0 {
		return -1
	}

	b := 1
	for b < len(h) && h[b] < n {
		b *= 2
	}

	l := b / 2
	r := b + 1
	if len(h)-1 < r {
		r = len(h) - 1
	}

	return loop(n, h, l, r)
}
