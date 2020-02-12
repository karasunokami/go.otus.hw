package list

type List struct {
	root Item
	len  int
}

type Item struct {
	Next, Prev *Item
	list       *List
	Value      interface{}
}

func (l *List) Init() *List {
	l.root.Next = &l.root
	l.root.Prev = &l.root
	l.len = 0
	return l
}

func (l *List) Len() int { return l.len }

func (l *List) First() *Item {
	if l.len == 0 {
		return nil
	}
	return l.root.Next
}

func (l *List) Last() *Item {
	if l.len == 0 {
		return nil
	}
	return l.root.Prev
}

func (l *List) insert(item, at *Item) *Item {
	next := at.Next
	at.Next = item
	item.Prev = at
	item.Next = next
	next.Prev = item
	item.list = l
	l.len++
	return item
}

func (l *List) remove(item *Item) *Item {
	item.Prev.Next = item.Next
	item.Next.Prev = item.Prev
	item.Next = nil
	item.Prev = nil
	item.list = nil
	l.len--
	return item
}

func (l *List) insertValue(v interface{}, at *Item) *Item {
	return l.insert(&Item{Value: v}, at)
}

func (l *List) PushFront(v interface{}) *Item {
	if l.root.Next == nil {
		l.Init()
	}

	return l.insertValue(v, &l.root)
}

func (l *List) PushBack(v interface{}) *Item {
	if l.root.Next == nil {
		l.Init()
	}

	return l.insertValue(v, l.root.Prev)
}

func (l *List) Remove(item *Item) interface{} {
	if item.list == l {
		l.remove(item)
	}
	return item.Value
}
