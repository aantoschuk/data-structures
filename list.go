package main

import "fmt"

type LinkedList[T any] struct {
	head *Node[T]
}

type operation[T any] func(node *Node[T], index int) bool

func (l *LinkedList[T]) Traverse(fn operation[T]) {
	current := l.head
	index := 0
	// Traverse until current is nil (end of the list)
	for current != nil {
		// Call the operation function on each node
		if fn(current, index) {
			return
		}
		// Move to the next node
		current = current.next
		index++
	}
}

func (l *LinkedList[T]) Insert(value T, position int) bool {
	// Create a new node
	node := &Node[T]{data: value}

	// If we need insert new head
	if position == 0 {
		node.next = l.head
		l.head = node
		return true
	}

	// Traverse through linked list and insert ad desirable position
	l.Traverse(func(n *Node[T], index int) bool {
		if index == position-1 {
			node.next = n.next
			n.next = node
			return true
		}
		return false
	})
	return false
}

// Stringer, fancy print list
func (l *LinkedList[T]) String() (result string) {
	// Traverse through the linked list and accumulate the node data into result
	l.Traverse(func(node *Node[T], index int) bool {
		// If this is not a head, then add arrow for a fancy view
		if index > 0 {
			result += " -> "
		}
		result += fmt.Sprintf("%v", node.data)
		// Return false is we need to iterave over each element in the list
		return false
	})

	return
}
