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
	front  *ListItem
	back   *ListItem
	length int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	l.length++

	val := ListItem{
		Value: v,
		Next:  l.front,
	}

	if l.front != nil {
		l.front.Prev = &val
		l.front = &val
	} else {
		l.front = &val
		l.back = &val
	}

	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.length++

	val := ListItem{
		Value: v,
		Prev:  l.back,
	}

	if l.back != nil {
		l.back.Next = &val
		l.back = &val
	} else {
		l.back = &val
		l.front = &val
	}

	return l.back
}

func (l *list) Remove(i *ListItem) {
	l.length--

	prev, next := i.Prev, i.Next
	if i.Next != nil && i.Prev != nil {
		prev.Next = i.Next
		next.Prev = i.Prev
		return
	}

	if next == nil && prev == nil {
		l.front = nil
		l.back = nil
	}

	if next != nil {
		l.front = next
		l.front.Prev = nil
	}

	if prev != nil {
		l.back = prev
		l.back.Next = nil
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	l.Remove(i)
	l.PushFront(i.Value)
}
