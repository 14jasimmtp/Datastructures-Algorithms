package main

type Tree struct{
	key int
	left *Tree
	right *Tree
}

type BST struct{
	root *Tree
}

func (t *BST) insert(key int){
	if t.root == nil{
		t.root = &Tree{key: key,left: nil,right: nil}
	}

	t.Insert(t.root,key)
	
}

func (t *BST) Insert(root *Tree,key int) *Tree{
	if root == nil{
		return &Tree{key: key,right: nil,left: nil} 
	}

	if key < root.key{
		root.left = t.Insert(root.left,key)
	}else{
		root.right = t.Insert(root.right,key)
	}

	return root
}

// func (t *BST) contains(key int) bool{
// 	if t.root
// }

func main(){
	tree:=BST{}

	tree.insert(5)
	tree.insert(4)
}