package bsearch

import (
	"reflect"
	"runtime"
	"sort"
	"strings"
)

// Algorithms available to use.
var Algorithms algorithms

// init initializes list of all available algorithms.
func init() {
	cf := []SearchFunc{
		Interpolation,
		Loop,
		LoopExponential,
		LoopSlicing,
		Recursive,
		RecursiveSlicing,
		Standard,
	}
	Algorithms = make(algorithms, len(cf))
	for _, f := range cf {
		Algorithms[f.Name()] = f
	}
}

// A Searcher searches for n (needle) in a sorted slice of ints h (haystack).
// The return value is the index of n or -1 if n is not present in h.
// The slice must be sorted in ascending order.
type Searcher interface {
	Search(n int, v []int) int
	Name() string
}

type algorithms map[string]Searcher

// Names returns sorted list of algorithms.
func (a algorithms) Names() []string {
	names := make([]string, 0, len(a))
	for k := range a {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}

//
func (a algorithms) Sorted() []Searcher {
	var sorted = make([]Searcher, 0, len(a))
	for _, name := range a.Names() {
		sorted = append(sorted, a[name])
	}

	return sorted
}

// String returns sorted and comma-space separated list of algorithms.
func (a algorithms) String() string {
	return strings.Join(a.Names(), ", ")
}

// The SearcherFunc type is an adapter to allow the use of
// ordinary functions as Searcher interface.
//
// If f is a function with the appropriate signature,
// SearcherFunc(f) is a Searcher that calls f.
type SearchFunc func(i int, v []int) int

// Search for n (needle) in a sorted slice of ints h (haystack).
// The return value is the index of n or -1 if n is not present in h.
// The slice must be sorted in ascending order.
func (f SearchFunc) Search(i int, v []int) int {
	return f(i, v)
}

func (f SearchFunc) Name() string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	field := strings.SplitAfter(name, ".")
	name = field[len(field)-1]
	name = strings.ToLower(name)
	return name
}
