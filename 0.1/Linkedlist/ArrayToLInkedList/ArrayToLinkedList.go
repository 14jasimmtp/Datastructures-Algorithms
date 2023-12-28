package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) ArrayToLinkedList(arr []int) {
	current := list.Head
	for _, n := range arr {
		newNode := &Node{Data: n, Next: nil}

		if list.Head == nil {
			list.Head = newNode
			current = newNode
		} else {
			current.Next = newNode
			current = current.Next
		}
	}
}

func (list *LinkedList) Display() {
	current := list.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	// Creating a linked list
	myList := LinkedList{}
	myList.ArrayToLinkedList([]int{1,3,4,6,7,3,4})
	myList.Display()
}
