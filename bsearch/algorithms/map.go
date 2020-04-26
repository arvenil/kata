package algorithms

// Map is a list of all available binary search algorithms indexed by name.
var Map map[string]Algorithm

// init Map from a Slice.
func init() {
	Map = Slice.Map()
}
