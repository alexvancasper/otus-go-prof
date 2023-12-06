package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	head *ListItem
	last *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	newHead := &ListItem{Value: v, Next: l.head, Prev: nil}
	if l.len == 0 {
		l.head = newHead
		l.last = newHead
	} else {
		l.head.Prev = newHead
		l.head = newHead
	}
	l.len++
	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	newBack := &ListItem{Value: v, Next: nil, Prev: l.last}
	if l.len == 0 {
		l.last = newBack
		l.head = newBack
	} else {
		l.last.Next = newBack
		l.last = newBack
	}
	l.len++
	return l.head
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}
	if (i == l.head || i == l.last) && l.len == 1 {
		l.head = nil
		l.last = nil
		l.len = 0
		return
	}
	if i == l.head {
		// first element
		l.head = l.head.Next
		l.head.Prev = nil
		l.len--
		return
	}
	if i == l.last {
		// last element
		prev := l.last.Prev
		l.last.Prev = nil
		prev.Next = nil
		l.last = prev
		l.len--
		return
	}
	prev := i.Prev
	next := i.Next
	prev.Next = next
	next.Prev = prev
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil || i == l.head {
		return
	}
	if i == l.last {
		prev := l.last.Prev
		prev.Next = nil
		l.last.Next = l.head
		l.head.Prev = l.last
		l.last.Prev = nil
		l.head = l.last
		l.last = prev
		return
	}

	prev := i.Prev
	next := i.Next
	prev.Next = next
	next.Prev = prev
	i.Next = l.head
	i.Prev = nil
	l.head.Prev = i
	l.head = i
}
