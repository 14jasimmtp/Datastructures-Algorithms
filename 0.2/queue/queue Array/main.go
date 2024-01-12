package main

import "fmt"

type Queue struct{
	data []int
}
func main(){
	Queue:=&Queue{}
	Queue.Enqueue(1)
	Queue.Enqueue(5)
	Queue.Enqueue(2)
	Queue.display()
	Queue.Dequeue()
	Queue.display()
}

func (s *Queue)Enqueue(val int){
	s.data = append(s.data, val)
}

func (s *Queue)Dequeue(){
	s.data = s.data[1:]
}

func (s *Queue)display(){
	for i:=0;i<len(s.data);i++{
		fmt.Printf("%d ->",s.data[i])
	}
	fmt.Println("nil")
}
