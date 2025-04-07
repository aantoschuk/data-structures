package main

import "fmt"

func main() {
	list := LinkedList[int]{head: &Node[int]{data: 1, next: nil}}
	list.Insert(2, 0)
	list.Insert(3, 2)
	// Since i pass l *LinkedList[T], i should call a pointer
	fmt.Println(&list)
}
