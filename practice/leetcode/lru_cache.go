package leetcode

// To implement an LRU (Least Recently Used) Cache in Go with
// time complexity for both Get and Put operations, you must combine a Hash Map (for fast lookups) and a Doubly Linked List (to maintain access order).
// Core Logic

//     Hash Map: Stores the key and a pointer to its corresponding Node in the linked list.
//     Doubly Linked List: Stores the key-value pairs. The Head represents the Most Recently Used (MRU) item, and the Tail represents the Least Recently Used (LRU) item.
//     Sentinel Nodes: Using dummy head and tail nodes simplifies the code by eliminating null checks during node removal or insertion

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node // Dummy sentinel nodes
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.moveToFront(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.value = value
		this.moveToFront(node)
	} else {
		newNode := &Node{key: key, value: value}
		this.cache[key] = newNode
		this.addNode(newNode)
		if len(this.cache) > this.capacity {
			lru := this.tail.prev
			this.removeNode(lru)
			delete(this.cache, lru.key)
		}
	}
}

// Helper: Adds node right after the dummy head (Most Recently Used)
func (this *LRUCache) addNode(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// Helper: Removes an existing node from the linked list
func (this *LRUCache) removeNode(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

// Helper: Moves a node to the front (MRU position)
func (this *LRUCache) moveToFront(node *Node) {
	this.removeNode(node)
	this.addNode(node)
}
