package main

import "fmt"

type Node struct{
	Key int
	Left *Node
	Right *Node
}
type Tree struct{
	Root *Node
}

func (tree *Tree) InOrder(root *Node){
	if root != nil{
		tree.InOrder(root.Left)
		fmt.Print(root.Key," ")
		tree.InOrder(root.Right)
	}
}

func (tree *Tree) PostOrder(root *Node){
	
	if root != nil{
		tree.PostOrder(root.Left)
		tree.PostOrder(root.Right)
		fmt.Print(root.Key," ")
	}
}

func (tree *Tree) PreOrder(root *Node){

	if root != nil {
		fmt.Print(root.Key," ")
		tree.PreOrder(root.Left)
		tree.PreOrder(root.Right)
	}
}

func (tree *Tree) Insert(key int){
	if tree.Root == nil{
		 tree.Root = &Node{Key: key,Left: nil,Right: nil}
	}else{
		tree.insert(tree.Root,key)
	}
}

func (tree *Tree) insert(root *Node,key int) *Node{
	if root == nil {
		return &Node{Key: key,Left: nil,Right: nil}
	}

	if key < root.Key{
		root.Left = tree.insert(root.Left,key)
	}else{
		root.Right = tree.insert(root.Right,key)
	}

	return root
}

func main(){
	tree:=Tree{}
	keys:=[]int{7,2,2,5,24,6,23,52,72,234,6,432,6,4}
	for _,key:=range keys{
		tree.Insert(key)
	}

	fmt.Print("Post-Order : ")
	tree.InOrder(tree.Root)
	fmt.Println()
	fmt.Print("In-Order : ")
	tree.PostOrder(tree.Root)
	fmt.Println()
	fmt.Print("Pre-Order : ")
	tree.PreOrder(tree.Root)
	fmt.Println()

}