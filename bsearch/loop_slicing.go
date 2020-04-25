package bsearch

func LoopSlicing(n int, v []int) (i int) {
	for len(v) > 0 {
		m := (len(v) - 1) / 2
		if v[m] < n {
			v = v[m+1:]
			i += m + 1
		} else if v[m] > n {
			v = v[:m]
		} else {
			return m + i
		}
	}

	// Return -1 when not found.
	return -1
}
