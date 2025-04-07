package main

// Generic Node which have any type
type Node[T any] struct {
	data T
	next *Node[T]
}
