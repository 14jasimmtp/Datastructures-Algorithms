package main

import "fmt"

const arraySize = 7

type HashTable struct {
	array [arraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key      string
	value    string
	nextNode *bucketNode
}

func (h *HashTable) Insert(key, value string) {
	index := hash(key)
	h.array[index].insert(key, value)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)

}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)

}

func (b *bucket) insert(k, value string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k, value: value}
		newNode.nextNode = b.head
		b.head = newNode
	} else {
		fmt.Println("already exist")
	}
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.nextNode

	}
	return false

}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.nextNode
		return
	}
	previousNode := b.head
	for previousNode != nil {
		if previousNode.nextNode.key == k {
			previousNode.nextNode = previousNode.nextNode.nextNode
		}
		previousNode = previousNode.nextNode
	}
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % arraySize
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	testHashtable := Init()

	testHashtable.Insert("django", "malayalam")

	fmt.Println(testHashtable.Search("django"))
	fmt.Println(testHashtable.Search("KYLE"))

}
