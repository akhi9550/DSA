package main

import "fmt"

//array size is the size of the hash table array
const arraySize = 7

//hash table which hold the array
type HashTable struct {
	array [arraySize]*bucket
}

//bucket structure-bucket is a linked list
type bucket struct {
	head *bucketNode
}

//bucket node structure-bucket node is a linkedlist node that holds the key and the address of the next node
type bucketNode struct {
	key      string
	nextNode *bucketNode
}

//insert will take a key and add it into the hashtable array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

//search will take the key and return true if it is in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)

}

//delete will take the key and e it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)

}

//insert will take in a key ,create a node with the key and insert the node in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.nextNode = b.head
		b.head = newNode
	} else {
		fmt.Println("already exist")
	}
}

//search will take the key and return true if the bucket has the key
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

//delete will take in a key and delete the node from the bucket
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

//hash function is to add the key into the hash table
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % arraySize
}

//init will create a bucket in each slot of the hash table ,it will return an address of an hashtable
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	testHashtable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}
	for _, v := range list {
		testHashtable.Insert(v)
	}
	testHashtable.Delete("STAN")
	fmt.Println(testHashtable.Search("STAN"))
	fmt.Println(testHashtable.Search("KYLE"))

}