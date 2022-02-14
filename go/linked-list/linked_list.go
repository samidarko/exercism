package linkedlist

import "fmt"

var ErrEmptyList = fmt.Errorf("emtpty list")

type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func NewList(args ...interface{}) *List {
	list := &List{}
	if len(args) == 0 {
		return list
	}
	node := &Node{Val: args[0]}
	list.first = node
	list.last = node
	for _, val := range args[1:] {
		last := list.Last()
		node = &Node{Val: val, prev: last}
		last.next = node
		list.last = node
	}
	return list
}

type List struct {
	first *Node
	last  *Node
}

func (l *List) PushFront(v interface{}) {
	first := l.First()
	node := &Node{Val: v, next: first}
	if first != nil {
		first.prev = node
	}
	if first == nil {
		l.last = node
	}
	l.first = node
}

func (l *List) PushBack(v interface{}) {
	last := l.Last()
	node := &Node{Val: v, prev: l.last}
	if last != nil {
		last.next = node
	}
	if last == nil {
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
	}
	if first.next == nil { // first is also last
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
	}
	if last.prev == nil { // last is also first
		l.first = nil
		l.last = nil
	}
	return last.Val, nil
}

func (l *List) Reverse() {
	node := l.First()
	for node != nil {
		next := node.Next()
		node.prev, node.next = node.next, node.prev
		node = next
	}
	l.first, l.last = l.last, l.first
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
