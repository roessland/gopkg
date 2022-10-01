package priorityqueue

import "container/heap"

type PriorityQueue[T any] struct {
	queue *queue[T]
}

func New[T any]() PriorityQueue[T] {
	s := make(queue[T], 0)
	return PriorityQueue[T]{
		queue: &s,
	}
}

func (pq *PriorityQueue[T]) Push(item T, priority float64) {
	it := newItem(item, priority)
	heap.Push(pq.queue, it)
}

// Pop removes the item in the heap with lowest priority, returning the item.
func (pq *PriorityQueue[T]) Pop() T {
	return heap.Pop(pq.queue).(item[T]).value
}

// PopPri removes the item in the heap with the lowest priority, returning
// both the item and the priority it had. This can be used for example in
// Dijkstra's algorithm if one adds duplicate items instead of decreasing
// the priority of existing items. If the popped priority is not the minimum
// distance, simply skip the item since we have already processed a copy of it
// with lower priority.
func (pq *PriorityQueue[T]) PopPri() (T, float64) {
	it := heap.Pop(pq.queue).(item[T])
	return it.value, it.priority
}

func (pq *PriorityQueue[T]) Len() int {
	return len(*pq.queue)
}

func newItem[T any](value T, priority float64) item[T] {
	return item[T]{
		value:    value,
		priority: priority,
	}
}

type item[T any] struct {
	value    T
	priority float64 // Distance from source to this node
}

type queue[T any] []item[T]

func (pq *queue[T]) Len() int { return len(*pq) }

func (pq *queue[T]) Less(i, j int) bool {
	return (*pq)[i].priority < (*pq)[j].priority
}

func (pq *queue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *queue[T]) Push(item_ any) {
	it := item_.(item[T])
	*pq = append(*pq, it)
}

func (pq *queue[T]) Pop() any {
	old := *pq
	n := len(old)
	it := old[n-1]
	*pq = old[0 : n-1]
	return it
}
