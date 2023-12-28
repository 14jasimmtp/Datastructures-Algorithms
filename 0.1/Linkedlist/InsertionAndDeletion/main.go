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
func (list *LinkedList) DisplayReverse() {
	fmt.Print("nil")
	list.PrintReverse(list.Head)
	fmt.Println()

}
func (list *LinkedList) PrintReverse(node *Node) {
	if node == nil {
		return
	}
	defer fmt.Printf("-> %d ", node.Data)

	list.PrintReverse(node.Next)
}

func (list *LinkedList) DeleteData(data int) {

	if list.Head == nil {
		return
	}

	if list.Head.Data == data {
		list.Head = list.Head.Next
		return
	}

	current := list.Head

	for current.Next != nil && current.Next.Data != data {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	}
}

func (list *LinkedList) InserNodeBeforeAndAfter(beforeData, AfterData, Data int) {
	current := list.Head
	newNodeAfter := &Node{Data: AfterData, Next: nil}

	for current != nil {
		if current.Data == Data {
			newNodeAfter.Next = current
		}
		current = current.Next
	}
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
	myList := LinkedList{}

	myList.InsertData(1)
	myList.InsertData(2)
	myList.InsertData(3)
	myList.InsertData(4)

	
	myList.InserNodeBeforeAndAfter(8, 9, 3)
	myList.Display()
	
}
