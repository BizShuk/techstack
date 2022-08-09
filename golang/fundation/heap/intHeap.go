package heap

import (
	"container/heap"
	"fmt"
)

// [Pattern]: [Go Heap] Standard IntHeap. Use negative value for MaxHeap without implement max heap
type IntHeap []int // An IntHeap is a min-heap of ints.

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func IntHeapMain() {
	h := &IntHeap{2, 1, 5} // [Notice]: [Go Heap] heap must be pointers due to struct itself will change what range of slice it points to.
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
