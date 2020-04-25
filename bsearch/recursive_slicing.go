package bsearch

func RecursiveSlicing(i int, v []int) int {
	m := len(v) / 2
	switch {
	case len(v) == 0:
		return -1 // not found
	case v[m] > i:
		return Recursive(i, v[:m])
	case v[m] < i:
		r := Recursive(i, v[m+1:])
		if r == -1 {
			return -1
		}
		return r + m + 1
	}
	return m
}
