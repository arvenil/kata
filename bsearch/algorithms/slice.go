package algorithms

import (
	"github.com/arvenil/kata/bsearch"
)

// Slice is a list of all available binary search algorithms.
// Register here any new Func.
var Slice = Algorithms{ //nolint:gochecknoglobals
	Func(bsearch.Interpolation),
	Func(bsearch.Loop),
	Func(bsearch.Exponential),
	Func(bsearch.LoopSlicing),
	Func(bsearch.Recursive),
	Func(bsearch.RecursiveSlicing),
	Func(bsearch.Standard),
}
