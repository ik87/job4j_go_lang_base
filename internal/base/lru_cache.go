package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	size   int
	length int
	Head   *Node
	Tail   *Node
}

func NewLruCache(size int) *LruCache {
	return &LruCache{
		size: size,
	}
}

func (l *LruCache) Put(key string, value string) {
	first := l.Head
	last := l.Tail
	for first != nil && last != nil {

		if first.Key == key {
			first.Value = value
			return
		}

		if last.Key == key {
			last.Value = value
			return
		}

		if first == last {
			break
		}

		first = first.Next
		last = last.Prev
	}

	newNode := Node{Key: key, Value: value}
	l.length++

	if l.Head == nil {
		l.Head = &newNode
		l.Tail = &newNode
		return
	}

	if l.length > l.size {
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
		l.length = l.size
	}

	l.Head.Prev = &newNode
	newNode.Next = l.Head
	l.Head = &newNode
}

func (l *LruCache) Get(key string) *string {

	if l.Head == nil {
		return nil
	}

	first := l.Head
	last := l.Tail
	tmpKey := l.Head.Key
	tmpValue := l.Head.Value

	for first != nil && last != nil {

		if first.Key == key {
			l.Head.Key = first.Key
			l.Head.Value = first.Value
			first.Key = tmpKey
			first.Value = tmpValue
			return &l.Head.Value
		}

		if last.Key == key {
			l.Head.Key = last.Key
			l.Head.Value = last.Value
			last.Key = tmpKey
			last.Value = tmpValue
			return &l.Head.Value
		}

		if first == last {
			return nil
		}

		first = first.Next
		last = last.Prev
	}

	return nil
}
