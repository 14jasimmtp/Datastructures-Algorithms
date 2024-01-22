package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}
type Tree struct {
	Root *Node
}

func (tree *Tree) InOrder(root *Node) {
	if root != nil {
		tree.InOrder(root.Left)
		fmt.Print(root.Key, " ")
		tree.InOrder(root.Right)
	}
}

func (tree *Tree) PostOrder(root *Node) {

	if root != nil {
		tree.PostOrder(root.Left)
		tree.PostOrder(root.Right)
		fmt.Print(root.Key, " ")
	}
}

func (tree *Tree) PreOrder(root *Node) {

	if root != nil {
		fmt.Print(root.Key, " ")
		tree.PreOrder(root.Left)
		tree.PreOrder(root.Right)
	}
}

func (tree *Tree) Insert(key int) {
	if tree.Root == nil {
		tree.Root = &Node{Key: key}
	} else {
		tree.insert(tree.Root, key)
	}
}

func (tree *Tree) insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{Key: key}
	}

	if key <= root.Key {
		root.Left = tree.insert(root.Left,key)
	} else {
		root.Right = tree.insert(root.Right,key)
	}

	return root
}

func (tree *Tree)BFS(root *Node){
	queue:=[]*Node{}
	queue = append(queue, root)

	for len(queue)>0{
		current:=queue[0]
		queue = queue[1:]

		fmt.Print(" ",current.Key)

		if current.Left != nil{
			queue =append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

func (tree *Tree) Delete(val int){
	Delete(tree.Root,val)
}
func Delete(root *Node,val int) *Node{
	if root == nil{
		return root
	}

	if val <root.Key{
		root.Left = Delete(root.Left,val)
	}else if val > root.Key{
		root.Right = Delete(root.Right,val)
	}else{
		if root.Left == nil && root.Right == nil{
			return nil
		}else if root.Left == nil{
			return root.Right
		}else if root.Right == nil{
			return root.Left
		}else{
			root.Key = min(root.Right)
			root.Right = Delete(root.Right,root.Key)
		}
	}

	return root
}

func min(root *Node) int{
	if root.Left == nil{
		return root.Key
	}

	for root.Left != nil{
		root =root.Left
	}

	return root.Key
}

func main() {
	tree := Tree{}
	keys := []int{7, 2, 2, 5, 24, 6, 23, 52, 72, 234, 6, 432, 6, 4}
	for _, key := range keys {
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
	tree.Delete(234)
	fmt.Print("BFS : ")
	tree.BFS(tree.Root)

}
