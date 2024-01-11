package main

import "fmt"

type Node struct{
	Data int
	Next *Node
}

type Stack struct{
	Top *Node
}

func(s *Stack) push(val int){
	newNode:=&Node{Data: val,Next: s.Top}
	s.Top = newNode
}

func (s *Stack) pop(){
	if s.Top == nil{
		fmt.Println("stack is empty.Can't pop")
		return
	}
	fmt.Printf("%d is popped from stack",s.Top.Data)
	fmt.Println()
	s.Top =s.Top.Next
}

func (s *Stack) display(){
	if s.Top == nil{
		fmt.Println("stack is empty")
		return
	}
	current:=s.Top
	for current != nil{
		fmt.Printf("%d->",current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func (s *Stack) minimum() int{
	min := s.Top.Data
	current :=s.Top

	for current != nil{
		if current.Data < min{
			min = current.Data
		}
		current = current.Next
	}
	return min 
}

func main(){
	stack:=&Stack{}

	stack.push(1)
	stack.push(8)
	stack.pop()
	stack.push(3)
	stack.push(7)

	stack.display()
	stack.push(4)
	fmt.Println(stack.minimum())
	stack.pop()
	stack.pop()
	stack.pop()
	stack.pop()
	stack.pop()


	stack.display()

}