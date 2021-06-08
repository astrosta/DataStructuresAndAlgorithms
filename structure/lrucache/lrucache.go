package lrucache

const (
	hostbit = uint64(^uint(0)) == ^uint64(0)
)

var cacheSize int

type LruNode struct {
	prev  *LruNode
	next  *LruNode
	hnext *LruNode
	key   int
	value interface{}
}

type LruList struct {
	head *LruNode
	tail *LruNode
}

type LruHashTable struct {
	node []LruNode
}

type LruCache struct {
	LruHashTable //快速定位数据
	LruList      //数据存储
	capacity     int
	used         int
}

func NewLruCache(size int) *LruCache {
	cacheSize = size
	return &LruCache{
		LruHashTable: LruHashTable{
			node: make([]LruNode, size),
		},
		LruList: LruList{
			head: nil,
			tail: nil,
		},
		capacity: size,
		used:     0,
	}
}

func (lc *LruCache) Get(key int) interface{} {
	node := lc.searchNode(key)
	if node == nil {
		return nil
	}

	lc.moveToTail(node)
	return node.value
}

func (lc *LruCache) Put(key int, value interface{}) {
	if node := lc.searchNode(key); node != nil {
		node.value = value
		lc.moveToTail(node)
	}
	lc.addNode(key, value)

	if lc.used > lc.capacity {
		lc.deleteNode()
	}

	return
}

func (lc *LruCache) addNode(key int, value interface{}) {
	newNode := &LruNode{
		key:   key,
		value: value,
	}
	tmp := lc.node[hash(key)]
	newNode.hnext = tmp.hnext
	tmp.hnext = newNode
	lc.used++

	if lc.tail == nil {
		lc.head, lc.tail = newNode, newNode
		return
	}

	lc.tail.next = newNode
	newNode.prev = lc.tail
	lc.tail = newNode
}

func (lc *LruCache) deleteNode() {
	if lc.head == nil {
		return
	}

	prev := &lc.node[hash(lc.head.key)]

	for prev.hnext != nil {
		if prev.hnext.key == lc.head.key {
			prev.hnext = prev.hnext.hnext
			lc.head = lc.head.next
			lc.head.prev = nil
			lc.used--
			break
		}

		prev = prev.hnext
	}

	return
}

func (lc *LruCache) searchNode(key int) *LruNode {
	if lc.tail == nil {
		return nil
	}

	temp := lc.node[hash(key)].hnext

	for temp != nil {
		if temp.key == key {
			return temp
		}

		temp = temp.hnext
	}

	return nil
}

func (lc *LruCache) moveToTail(node *LruNode) {
	if lc.tail == node {
		return
	}

	if lc.head == node {
		lc.head = node.next
		lc.head.prev = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	node.next = nil
	lc.tail.next = node
	node.prev = lc.tail
	lc.tail = node
}

func hash(key int) int {
	if hostbit {
		return (key ^ (key >> 32)) & (cacheSize - 1)
	}

	return (key ^ (key >> 16)) & (cacheSize - 1)
}
