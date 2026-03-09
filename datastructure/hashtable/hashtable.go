package hashtable

import "fmt"

const arraySize int = 7

type hashTable struct {
	array [arraySize]*bucket
}

func NewHashTable() *hashTable {
	hashTable := &hashTable{}
	for i := range hashTable.array {
		hashTable.array[i] = &bucket{}
	}

	return hashTable
}

func (h *hashTable) Insert(key string) {
	if !h.Search(key) {
		index := hash(key)
		h.array[index].insert(key)
	} else {
		fmt.Println("already exist")
	}
}

func (h *hashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *hashTable) Delete(key string) bool {
	index := hash(key)
	return h.array[index].delete(key)
}

type bucketNode struct {
	key  string
	next *bucketNode
}

type bucket struct {
	head *bucketNode
}

func (b *bucket) insert(k string) {
	newNode := &bucketNode{key: k, next: b.head}
	b.head = newNode
}

func (b *bucket) search(k string) bool {
	cn := b.head
	for cn != nil {
		if cn.key == k {
			return true
		}

		cn = cn.next
	}

	return false
}

func (b *bucket) delete(k string) bool {
	if b.head.key == k {
		b.head = b.head.next
		return true
	}

	prev := b.head
	for prev.next != nil {
		if prev.next.key == k {
			prev.next = prev.next.next
			return true
		}

		prev = prev.next
	}

	return false
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % arraySize
}
