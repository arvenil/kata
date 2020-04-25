package bsearch

func Interpolation(i int, v []int) int {
	if len(v) == 0 {
		return -1
	}

	l := 0
	r := len(v) - 1
	for (v[r] != v[l]) && (i >= v[l]) && (i <= v[r]) {
		m := l + (i-v[l])*(r-l)/(v[r]-v[l])

		if v[m] < i {
			l = m + 1
		} else if v[m] > i {
			r = m - 1
		} else {
			return m
		}
	}

	if v[l] == i {
		return l
	}

	return -1
}
