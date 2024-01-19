package main

import "fmt"

const size = 26

type Node struct {
	Children [size]*Node
	IsEnd    bool
}

type Trie struct {
	Root *Node
}

func (t *Trie) insert(s string) {
	current := t.Root
	for i:=0;i<len(s);i++ {
		charIdx := s[i] - 'a'
		if current.Children[charIdx] == nil {
			current.Children[charIdx] = &Node{}
		}
		current = current.Children[charIdx]
	}
	current.IsEnd = true
}

func (t *Trie) search(s string) bool {
	current := t.Root

	for i := 0;i<len(s);i++ {
		charIdx := s[i] - 'a'
		if current.Children[charIdx] == nil {
			return false
		}
		current = current.Children[charIdx]
	}

	return current.IsEnd
}

func main() {
	t := Trie{
		Root: &Node{},
	}
	str := []string{
		"jasim",
		"jaseel",
		"jawhara",
	}

	for _, strs := range str {
		t.insert(strs)
	}

	fmt.Println(t.search("jawhara"))
}
