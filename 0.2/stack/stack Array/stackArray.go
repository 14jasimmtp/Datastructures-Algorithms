package main

import "fmt"

type Stack struct{
	data []int
	
}
func main(){
	stack:=&Stack{}
	stack.push(1)
	stack.push(5)
	stack.push(2)
	stack.display()
	stack.pop()
	stack.display()
}

func (s *Stack)push(val int){
	s.data = append(s.data, val)
}

func (s *Stack)pop(){
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack)display(){
	for i:=len(s.data)-1;i>=0;i--{
		fmt.Printf("%d ->",s.data[i])
	}
	fmt.Println("nil")
}