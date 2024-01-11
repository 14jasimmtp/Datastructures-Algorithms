package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type Queue struct {
	Front *Node
	Rear  *Node
}

func (q *Queue) Enqueue(val int) {
	newNode := &Node{Val: val, Next: nil}
	if q.Rear == nil {
		q.Rear = newNode
		q.Front = newNode
	} else {
		q.Rear.Next = newNode
		q.Rear = q.Rear.Next
	}
}

func (q *Queue) Dequeue() {
	if q.Front == nil {
		return
	}

	q.Front = q.Front.Next
	if q.Front == nil {
		q.Rear = nil
	}
}

func (q *Queue) Display() {
	if q.Front == nil {
		fmt.Println("stack is empty")
		return
	}
	current := q.Front
	for current != nil {
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(5)
	queue.Enqueue(13)
	queue.Enqueue(4)
	queue.Dequeue()
	queue.Display()

}
