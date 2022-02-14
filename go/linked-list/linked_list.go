package linkedlist

import "fmt"

var ErrEmptyList = fmt.Errorf("emtpty list")

type Node struct {
	Val        interface{}
	next, prev *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func NewList(args ...interface{}) *List {
	list := new(List)
	for _, val := range args {
		list.PushBack(val)
	}
	return list
}

type List struct {
	first, last *Node
}

func (l *List) PushFront(v interface{}) {
	first := l.First()
	node := &Node{Val: v, next: first}
	if first != nil {
		first.prev = node
	} else {
		l.last = node
	}
	l.first = node
}

func (l *List) PushBack(v interface{}) {
	last := l.Last()
	node := &Node{Val: v, prev: last}
	if last != nil {
		last.next = node
	} else {
		l.first = node
	}
	l.last = node
}

func (l *List) PopFront() (interface{}, error) {
	first := l.First()
	if first == nil {
		return nil, ErrEmptyList
	}
	if first.next != nil {
		l.first = first.next
		l.first.prev = nil
	} else {
		// no more elements in the list
		l.first = nil
		l.last = nil
	}
	return first.Val, nil
}

func (l *List) PopBack() (interface{}, error) {
	last := l.Last()
	if last == nil {
		return nil, ErrEmptyList
	}
	if last.prev != nil {
		l.last = last.prev
		l.last.next = nil
	} else {
		// no more elements in the list
		l.first = nil
		l.last = nil
	}
	return last.Val, nil
}

func (l *List) Reverse() {
	for node := l.First(); node != nil; node = node.Prev() {
		node.prev, node.next = node.next, node.prev
	}
	l.first, l.last = l.last, l.first
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
