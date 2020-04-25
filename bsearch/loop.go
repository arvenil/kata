package bsearch

func Loop(i int, v []int) int {
	l := 0
	r := len(v) - 1
	return loop(i, v, l, r)
}

func loop(i int, v []int, l, r int) int {
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
