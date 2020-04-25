package bsearch

func LoopExponential(i int, v []int) int {
	if len(v) == 0 {
		return -1
	}

	b := 1
	for b < len(v) && v[b] < i {
		b *= 2
	}

	l := b / 2
	r := b + 1
	if len(v)-1 < r {
		r = len(v) - 1
	}

	return loop(i, v, l, r)
}
