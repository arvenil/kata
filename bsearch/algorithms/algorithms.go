package algorithms

import (
	"sort"
	"strings"
)

// Algorithms type enables bulk operations on slice of Algorithms.
type Algorithms []Algorithm

// Map of Algorithms indexed by name.
func (f Algorithms) Map() map[string]Algorithm {
	var funcs = make(map[string]Algorithm, len(f))
	for _, f := range f {
		funcs[f.Name()] = f
	}
	return funcs
}

// Names of Algorithms as sorted slice of strings.
func (f Algorithms) Names() []string {
	names := make([]string, 0, len(f))
	for _, f := range f {
		names = append(names, f.Name())
	}
	sort.Strings(names)
	return names
}

// String returns sorted and comma-space-separated Algorithm names.
func (f Algorithms) String() string {
	return strings.Join(f.Names(), ", ")
}
