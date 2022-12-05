package stackqueue

// Stack is LIFO queue data structure.
type Stack[T any] []T

// New creates a new stack.
func New[T any]() Stack[T] {
	s := make([]T, 0)
	return s
}

// Push adds an item to the stack.
func (stack *Stack[T]) Push(item T) {
	(*stack) = append(*stack, item)
}

// Pop remove the item from the top of the stack and returns it.
func (stack *Stack[T]) Pop() T {
	item := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return item
}

// Peek returns the item at the top of the stack.
func (stack *Stack[T]) Peek() T {
	return (*stack)[len(*stack)-1]
}

// Len returns the number of items in the stack.
func (stack *Stack[T]) Len() int {
	return len(*stack)
}
