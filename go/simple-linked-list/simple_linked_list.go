package linkedlist

import "errors"

// List type
type List struct {
	head *Element
	size uint
}

// Element type
type Element struct {
	value int
	next  *Element
}

// New returns a new List
func New(elements []int) *List {
	list := new(List)
	for _, element := range elements {
		list.Push(element)
	}
	return list
}

// Push add a new element to List
func (l *List) Push(value int) {
	element := &Element{
		value: value,
		next:  l.head,
	}
	l.head = element
	l.size++
}

// Size returns the List size
func (l *List) Size() uint {
	return l.size
}

// Pop returns the last inserted Element
func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("cannot pop from empty list")
	}
	value := l.head.value
	l.head = l.head.next
	l.size--
	return value, nil
}

// Array returns an array of elements
func (l *List) Array() []int {
	array := make([]int, 0)
	for element := l.head; element != nil; element = element.next {
		array = append([]int{element.value}, array...)
	}
	return array
}

// Reverse return a reversed list of elements
func (l *List) Reverse() *List {
	var prev, curr, next *Element
	curr = l.head
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev, curr = curr, next
	}
	l.head = prev
	return l
}
