package bsearch

import (
	"sort"
)

func Standard(i int, v []int) int {
	n := sort.SearchInts(v, i)
	if n < len(v) && v[n] == i {
		return n
	}
	return -1
}
