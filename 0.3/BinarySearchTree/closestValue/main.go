package main

import (
	"fmt"
	"math"
)

type Tree struct {
	key   int
	left  *Tree
	right *Tree
}

type BST struct {
	root *Tree
}

func (t *BST) insert(key int) {
	if t.root == nil {
		t.root = &Tree{key: key, left: nil, right: nil}
	}

	t.Insert(t.root, key)

}

func (t *BST) Insert(root *Tree, key int) *Tree {
	if root == nil {
		return &Tree{key: key, right: nil, left: nil}
	}

	if key < root.key {
		root.left = t.Insert(root.left, key)
	} else {
		root.right = t.Insert(root.right, key)
	}

	return root
}

func (t *BST) Closest(root *Tree, key float64) float64 {
	closest := float64(t.root.key)
	current := t.root

	for current != nil {
		if math.Abs((float64(current.key) - key)) < math.Abs((closest - key)) {
			closest = float64(current.key)
		}

		if key < float64(current.key) {
			current = current.left
		} else if key > float64(current.key) {
			current = current.right
		} else {
			break
		}
	}
	return closest
}

// func (t *BST) contains(key int) bool{
// 	if t.root
// }

func main() {
	tree := BST{}

	tree.insert(5)
	tree.insert(4)
	tree.insert(7)
	fmt.Println(tree.Closest(tree.root, 6))

}
