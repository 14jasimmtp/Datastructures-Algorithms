package main

import "fmt"


type Node struct {
	Children map[string]*Node
	IsEnd    bool
}

type Trie struct {
	Root *Node
}

func (t *Trie) insert(s string) {
	current := t.Root
	for i:=0;i<len(s);i++ {

		char:= string(s[i])

		if current.Children[char] == nil {

			current.Children[char] = &Node{Children: make(map[string]*Node)}

		}
		current = current.Children[char]
	}
	current.IsEnd = true
}

func (t *Trie) search(s string) bool {
	current := t.Root

	for i := 0;i<len(s);i++ {
		char := string(s[i])
		if current.Children[char] == nil {
			return false
		}
		current = current.Children[char]
	}

	return current.IsEnd
}

func (t *Trie) suggestions(prefix string) []string {
	word := []string{}

	current := t.Root

	for i := 0; i < len(prefix); i++ {
		char := string(prefix[i])

		if current.Children[char] == nil {
			return []string{}
		}

		current = current.Children[char]
	}

	t.containsPrefix(prefix, &word, current)

	return word
}

func (t *Trie) containsPrefix(prefix string, word *[]string, current *Node) {
	if current.IsEnd {
		*word = append(*word, prefix)
	}

	for char, child := range current.Children {
		t.containsPrefix(prefix+char, word, child)
	}
}


func main() {
	t := Trie{
		Root: &Node{
			Children: make(map[string]*Node),
		},
	}
	str := []string{
		"jasim",
		"jaseel",
		"jawhara",
		"jseel",
		"f",
	}

	for _, strs := range str {
		t.insert(strs)
	}

	fmt.Println(t.search("jawhara"))

	fmt.Println(t.suggestions("k"))
}
