package main

//使用双向链表
type node struct {
	key, value int
	prev, next *node
}

func initNode(key, value int) *node {
	return &node{
		key:   key,
		value: value,
	}
}

//LRUCache 缓存
type LRUCache struct {
	capacity   int
	size       int
	contentMap map[int]*node
	//添加一个假的头部节点跟尾部节点，选择的时候比较好判断
	head, tail *node
}

//ConstructorLRUCache 初始化
func ConstructorLRUCache(capacity int) LRUCache {
	contentMap := make(map[int]*node, capacity)

	l := LRUCache{
		size:       0,
		capacity:   capacity,
		contentMap: contentMap,
		head:       initNode(0, 0),
		tail:       initNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

//Get 获取值
func (l *LRUCache) Get(key int) int {
	if _, ok := l.contentMap[key]; !ok {
		return -1
	}
	node := l.contentMap[key]
	l.moveHead(node)
	return node.value
}

//Put 存放值
func (l *LRUCache) Put(key int, value int) {
	if _, ok := l.contentMap[key]; !ok {
		node := initNode(key, value)
		l.contentMap[key] = node
		l.addToHead(node)
		l.size++
		if l.size > l.capacity {
			removed := l.removeTail()
			delete(l.contentMap, removed.key)
			l.size--
		}
	} else {
		node := l.contentMap[key]
		node.value = value
		l.moveHead(node)
	}

}

//移动到头部
func (l *LRUCache) moveHead(n *node) {
	l.removeNode(n)
	l.addToHead(n)
}

//removeNode 当前节点移除
func (l *LRUCache) removeNode(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (l *LRUCache) addToHead(n *node) {
	//插入进去
	n.prev = l.head
	n.next = l.head.next
	//第三个节点连上
	l.head.next.prev = n
	//第一个节点连上
	l.head.next = n
}

//添加到尾部
func (l *LRUCache) removeTail() *node {
	node := l.tail.prev
	l.removeNode(node)
	return node
}
