package algorithms

import (
	"reflect"
	"runtime"
	"strings"
)

// The Func type is an adapter to allow the use of
// ordinary functions as an Algorithm interface.
//
// If f is a function with the appropriate signature,
// Func(f) is an Algorithm that calls f.
type Func func(n int, h []int) int

// Search for n (needle) in a sorted slice of ints h (haystack).
// The return value is the index of n or -1 if n is not present in h.
// The slice must be sorted in ascending order.
func (f Func) Search(n int, h []int) int {
	return f(n, h)
}

// Name of the algorithm.
func (f Func) Name() string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	field := strings.SplitAfter(name, ".")
	name = field[len(field)-1]
	name = strings.ToLower(name)
	return name
}

// Verify that Func fulfils Algorithm interface.
var _ Algorithm = Func(func(i int, v []int) int { return 0 })
