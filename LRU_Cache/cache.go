package lrucache

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

type LRUCache struct {
	Capacity int
	Cache    map[int]*Node
	Head     *Node
	Tail     *Node
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		Capacity: cap,
		Cache: make(map[int]*Node),
	}
}

func (l *LRUCache) Get(key int) int{
	if node,ok := l.Cache[key]; ok {
		l.remove(node)
		l.addToFront(node)
		return node.Value
	}
	return -1
}

func (l *LRUCache) Put(key,value int) {
	if node,ok := l.Cache[key]; ok {
		node.Value = value
		l.remove(node)
		l.addToFront(node)
	}else{
		if len(l.Cache) >= l.Capacity {
			delete(l.Cache,l.Tail.Key)
			l.remove(l.Tail)
		}
		newNode := &Node{Key: key,Value: value}
		l.Cache[key] = newNode
		l.addToFront(newNode)
	} 
}


func (l *LRUCache) addToFront(node *Node) {
	node.Next = l.Head
	node.Prev = nil 
	if l.Head != nil {
		l.Head.Prev = node
	}
	l.Head = node
	if l.Tail == nil {
		l.Tail = node
	}
}

func (l *LRUCache) remove(node *Node) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}else{
		l.Head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}else {
		l.Tail = node.Prev
	}
}