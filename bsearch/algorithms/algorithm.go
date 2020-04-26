package algorithms

// An Algorithm searches for n (needle) in a sorted slice of ints h (haystack).
// The return value is the index of n or -1 if n is not present in h.
// The slice must be sorted in ascending order.
// Algorithm has name.
type Algorithm interface {
	Search(n int, h []int) int
	Name() string
}
