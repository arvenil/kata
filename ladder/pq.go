package ladder

// A node in a priorityQueue.
type node struct {
	word   *word
	fScore int
	gScore int
	prev   *node // Previous node.
	// The index is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A priorityQueue implements heap.Interface and holds slice of *node.
type priorityQueue []*node

// Len is the number of elements in the collection.
func (pq priorityQueue) Len() int { return len(pq) }

// Less reports whether the element with
// index i should sort before the element with index j.
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].gScore < pq[j].gScore
}

// Swap swaps the elements with indexes i and j.
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// push pushes the element x into priorityQueue.
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*node)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the highest priority element from the priorityQueue.
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
