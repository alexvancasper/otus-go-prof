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
	back *ListItem
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
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newHead := &ListItem{Value: v, Next: l.head, Prev: nil}
	if l.len == 0 {
		l.head = newHead
		l.back = newHead
	} else {
		l.head.Prev = newHead
		l.head = newHead
	}
	l.len++
	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	newBack := &ListItem{Value: v, Next: nil, Prev: l.back}
	if l.len == 0 {
		l.back = newBack
		l.head = newBack
	} else {
		l.back.Next = newBack
		l.back = newBack
	}
	l.len++
	return l.head
}

func (l *list) Remove(i *ListItem) {
	if i == l.head {
		l.head = l.head.Next
		l.head.Prev = nil
		l.len--
		return
	}

	if i == l.back {
		l.back = l.back.Prev
		l.back.Next = nil
		l.len--
		return
	}

	cur := l.head
	prev := cur.Prev
	next := cur.Next
	for cur != nil {
		if cur.Value == i.Value {
			prev.Next = cur.Next
			next.Prev = prev
			l.len--
			return
		}
		cur = cur.Next
		prev = cur.Prev
		next = cur.Next
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.head {
		return
	}

	if i == l.back {
		prev := l.back.Prev
		l.back.Next = l.head
		l.head.Prev = l.back
		l.back.Prev = nil
		prev.Next = nil
		l.head = l.back
		l.back = prev
		return
	}

	cur := l.head
	prev := cur.Prev
	next := cur.Next
	for cur != nil {
		if cur == i {
			prev.Next = cur.Next
			next.Prev = cur.Prev
			cur.Next = l.head
			cur.Prev = nil
			l.head = cur
			return
		}
		cur = cur.Next
		prev = cur.Prev
		next = cur.Next
	}
}
