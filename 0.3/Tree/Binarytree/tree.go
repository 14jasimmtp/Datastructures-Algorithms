package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (tree *BinaryTree) Insert(key int) {
	if tree.Root == nil {
		tree.Root = &Node{Key: key, Left: nil, Right: nil}
	} else {
		tree.insert(tree.Root, key)
	}
}

func (tree *BinaryTree) insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{Key: key, Left: nil, Right: nil}
	}

	if key < root.Key {
		root.Left = tree.insert(root.Left, key)
	} else {
		root.Right = tree.insert(root.Right, key)
	}

	return root

}

func (tree *BinaryTree) Delete(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key < root.Key {
		root.Left = tree.Delete(root.Left, key)
	} else if key > root.Key {
		root.Right = tree.Delete(root.Right, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		root.Key = tree.min(root.Right)
		root.Right = tree.Delete(root.Right, root.Key)

	}
	return root
}

func (tree *BinaryTree) min(root *Node) int {
	if root.Left == nil {
		return root.Key
	}

	Key := tree.min(root.Left)

	return Key
}

func (tree *BinaryTree) max(root *Node) int {
	if root.Right == nil {
		return root.Key
	}

	Key := tree.max(root.Right)

	return Key
}

func (tree *BinaryTree) DisplayInOrder(root *Node) {
	if root != nil {
		tree.DisplayInOrder(root.Left)
		fmt.Print(root.Key, " ")
		tree.DisplayInOrder(root.Right)
	}
}

func main() {

	tree := BinaryTree{}
	keys := []int{5, 3, 8, 4, 6, 34, 72, 3, 8, 4, 7}

	for _, key := range keys {
		tree.Insert(key)
	}
	tree.Delete(tree.Root, 8)
	tree.Delete(tree.Root, 8)
	tree.DisplayInOrder(tree.Root)

	fmt.Println("max value : ", tree.max(tree.Root))
	fmt.Println("min value : ", tree.min(tree.Root))
}
