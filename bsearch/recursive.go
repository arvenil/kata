package bsearch

func Recursive(i int, v []int) int {
	return recursive(i, v, 0, len(v)-1)
}

func recursive(i int, v []int, l, h int) int {
	if l > h {
		return -1
	}
	m := (l + h) / 2
	if v[m] == i {
		return m
	}
	if i < v[m] {
		return recursive(i, v, l, m-1)
	}
	return recursive(i, v, m+1, h)
}
