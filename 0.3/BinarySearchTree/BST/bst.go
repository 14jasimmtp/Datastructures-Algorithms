package main

import "fmt"

type Tree struct {
	key   int
	left  *Tree
	right *Tree
}

type BST struct {
	root *Tree
}


// func (n *BST) Insert(value int) *Tree {
// 	if n == nil {
// 		return &Tree{key: value}
// 	}

// 	if value <= n.key {
// 		n.root.left = n.Insert()
// 	} else {
// 		n.right = n.right.Insert(value)
// 	}
// 	return n
// }


func (tree *BST) Insert(key int) {
	if tree.root == nil {
		tree.root = &Tree{key: key}
	} else {
		insert(tree.root, key)
	}
}

func insert(root *Tree, key int) *Tree {
	if root == nil {
		return &Tree{key: key}
	}

	if key < root.key {
		root.left = insert(root.left, key)
	} else {
		root.right = insert(root.right, key)
	}

	return root

}

func BFS(root *Tree) {
	queue := []*Tree{}

	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		current := queue[0]
		fmt.Print(current.key, " ")
		queue = queue[1:]

		if current.left != nil {
			queue = append(queue, current.left)
		}

		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
}



func main() {
	tree := BST{}
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(9)
	BFS(tree.root)
}
