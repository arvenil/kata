package algorithms

import (
	"sort"
	"strings"
)

// Algorithms type enables bulk operations on slice of Algorithms.
type Algorithms []Algorithm

// Map of Algorithms indexed by name.
func (a Algorithms) Map() map[string]Algorithm {
	fs := make(map[string]Algorithm, len(a))
	for _, f := range a {
		fs[f.String()] = f
	}

	return fs
}

// Names of Algorithms as sorted slice of strings.
func (a Algorithms) Names() []string {
	names := make([]string, 0, len(a))
	for _, f := range a {
		names = append(names, f.String())
	}

	sort.Strings(names)

	return names
}

// String returns sorted and comma-space-separated Algorithm names.
func (a Algorithms) String() string {
	return strings.Join(a.Names(), ", ")
}
