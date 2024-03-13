package dsa

type DoubleListNode struct {
	Val  int
	Next *DoubleListNode
	Prev *DoubleListNode
}

type LRUCache struct {
	capacity int
	h        map[int]int
	head     *DoubleListNode
	tail     *DoubleListNode
	nodeMap  map[int]*DoubleListNode
	l        int
}

func ConstructorLru(capacity int) LRUCache {
	return LRUCache{capacity: capacity, h: make(map[int]int), head: nil, tail: nil, nodeMap: make(map[int]*DoubleListNode), l: 0}
}

func (this *LRUCache) Get(key int) int {
	// fmt.Println("get: ", key, " ", this)
	if val, ok := this.h[key]; !ok || val == -1 {
		return -1
	} else {
		this.bringNodeToFront(key)
		return val
	}
}

func (this *LRUCache) Put(key int, value int) {
	// fmt.Println("put: ", key, " ", value, " ", this)
	// if the value getting inserted is not already present in the cache
	// add a new entry at tail and remove the oldest entry i.e the head
	if val, ok := this.h[key]; !ok || val == -1 {
		if this.l < this.capacity {
			this.h[key] = value
			this.l++
			// if its the first entry then just create the default entry as head and tal
			if this.head == nil {
				this.head = &DoubleListNode{Val: key, Next: nil, Prev: nil}
				this.tail = this.head
			} else { // if its not the first entry in cache with space add the entry at the end
				this.tail.Next = &DoubleListNode{Val: key, Next: nil, Prev: this.tail}
				this.tail = this.tail.Next
			}
			this.nodeMap[key] = this.tail
			return
		}
		// add new entry at the tail
		this.h[key] = value
		this.tail.Next = &DoubleListNode{Val: key, Next: nil, Prev: this.tail}
		this.tail = this.tail.Next
		this.nodeMap[key] = this.tail
		// remove the head
		this.h[this.head.Val] = -1
		this.nodeMap[this.head.Val] = nil
		this.head.Next.Prev = nil
		this.head = this.head.Next
		// if space is available in the cache
	} else {
		// fmt.Println("hola")
		this.bringNodeToFront(key)
		this.h[key] = value
	}
}

func (this *LRUCache) bringNodeToFront(key int) {
	// fmt.Println("brining key to the front ", key)
	// if the inserted value is already present in the cache
	// then make it the most recently accessed element
	node := this.nodeMap[key]

	if node == this.tail {
		return
	}
	// if its the oldest element in the cache then make sure to keep the prev as nil
	next := node.Next
	prev := node.Prev

	if next != nil {
		next.Prev = prev
	}
	if prev != nil {
		prev.Next = next
	} else {
		this.head = next
	}

	node.Prev = this.tail
	node.Next = nil
	this.tail.Next = node
	this.tail = this.tail.Next
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
