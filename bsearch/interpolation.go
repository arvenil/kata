package bsearch

// Interpolation searches for n (needle) in a sorted slice of ints h (haystack).
// The return value is the index of n or -1 if n is not present in h.
// The slice must be sorted in ascending order.
//
// Interpolation algorithms estimates the position of the target value,
// taking into account the lowest and highest elements in the array as well as length of the array.
// It works on the basis that the midpoint is not the best guess in many cases.
// For example, if the target value is close to the highest element in the array,
// it is likely to be located near the end of the array.
//
// https://en.wikipedia.org/wiki/Binary_search_algorithm#Interpolation_search
func Interpolation(n int, h []int) int {
	if len(h) == 0 {
		return -1
	}

	l := 0
	r := len(h) - 1
	for (h[r] != h[l]) && (n >= h[l]) && (n <= h[r]) {
		m := l + (n-h[l])*(r-l)/(h[r]-h[l])

		if h[m] < n {
			l = m + 1
		} else if h[m] > n {
			r = m - 1
		} else {
			return m
		}
	}

	if h[l] == n {
		return l
	}

	return -1
}
