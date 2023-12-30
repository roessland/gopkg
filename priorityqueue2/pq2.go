// Package priorityqueue2 proves a priority queue data structure. It uses
// generics for both the item and the priority.
// The lowest priority is popped first.
//
// Priority types: cmp.Ordered.
// - float32/float64: Min is -Inf, max is +Inf. NaN is not allowed!
// - string: Min is "". Sorting is alphanumerical. E.g. "" < "44" < "5"
// - uint(n): Min is 0, max is math.MaxUint(n).
// - int(n): Min is math.MinInt(n), max is math.MaxInt(n).
package priorityqueue2

import (
	"cmp"
	"container/heap"
)

type PriorityQueue[T any, P cmp.Ordered] struct {
	heap arrayHeap[T, P]
}

// New creates an empty priority queue.
func New[T any, P cmp.Ordered]() *PriorityQueue[T, P] {
	arr := arrayHeap[T, P](make([]heapItem[T, P], 0))
	return &PriorityQueue[T, P]{
		heap: arr,
	}
}

// Push adds an item to the priority queue with the given priority.
func (pq *PriorityQueue[T, P]) Push(item T, pri P) {
	hItem := heapItem[T, P]{item: item, priority: pri}
	heap.Push(&pq.heap, hItem)
}

// Pop removes the item in the heap with lowest priority, returning the item.
func (pq *PriorityQueue[T, P]) Pop() T {
	return heap.Pop(&pq.heap).(heapItem[T, P]).item
}

// PopPri removes the item in the heap with the lowest priority, returning
// both the item and the priority it had. This can be used for example in
// Dijkstra's algorithm if one adds duplicate items instead of decreasing
// the priority of existing items. If the popped priority is not the minimum
// distance, simply skip the item since we have already processed a copy of it
// with lower priority.
func (pq *PriorityQueue[T, P]) PopPri() (T, P) {
	it := heap.Pop(&pq.heap).(heapItem[T, P])
	return it.item, it.priority
}

// Len returns the number of items in the priority queue.
func (pq *PriorityQueue[T, P]) Len() int {
	return len(pq.heap)
}

// arrayHeap implements heap.Interface.
type arrayHeap[T any, P cmp.Ordered] ([]heapItem[T, P])

type heapItem[T any, P cmp.Ordered] struct {
	item     T
	priority P
}

var _ heap.Interface = new(arrayHeap[int, float64])

func (ah arrayHeap[T, P]) Len() int {
	return len(ah)
}

func (ah arrayHeap[T, P]) Less(i int, j int) bool {
	return ah[i].priority < ah[j].priority
}

func (ah *arrayHeap[T, P]) Pop() any {
	hItem := (*ah)[len(*ah)-1]
	*ah = (*ah)[:len(*ah)-1]
	return hItem
}

func (ah *arrayHeap[T, P]) Push(x any) {
	*ah = append(*ah, x.(heapItem[T, P]))
}

func (ah arrayHeap[T, P]) Swap(i int, j int) {
	ah[i], ah[j] = ah[j], ah[i]
}
