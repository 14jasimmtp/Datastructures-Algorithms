package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) InsertData(data int) {
	newNode := &Node{Data: data, Next: nil}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func (list *LinkedList) RemoveDuplicates() {
	seen := make(map[int]bool)
	var prev *Node

	current := list.Head
	for current != nil {
		if seen[current.Data] {
			prev.Next = current.Next
		} else {
			seen[current.Data] = true
			prev = current
		}
		current = current.Next
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

	// Inserting elements into the linked list
	myList.InsertData(1)
	myList.InsertData(4)
	myList.InsertData(2)
	myList.InsertData(3)
	myList.InsertData(4)
	myList.InsertData(4)
	myList.InsertData(3)

	myList.RemoveDuplicates()
	myList.Display()
}