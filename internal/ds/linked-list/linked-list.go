package linked_list

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	len  int
	size int
	head *Node
}

type Node struct {
	Value interface{}
	next  *Node
}

func NewLinkedList(size int) LinkedList {
	return LinkedList{
		size: size,
	}
}

func (l *LinkedList) Add(v interface{}) error {
	if l.size == l.len {
		return errors.New("list is full")
	}

	if l.len == 0 {
		l.head = &Node{Value: v}
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}

		current.next = &Node{Value: v}
	}

	l.len++

	return nil
}

func (l *LinkedList) InsertBefore(v, el interface{}) error {
	if l.size == l.len {
		return errors.New("list is full")
	}

	if l.len == 0 {
		return errors.New("list is empty")
	}

	current := l.head

	if current.Value == el {
		l.head = &Node{Value: v, next: current}
		l.len++
		return nil
	}

	for current != nil {
		if current.next == nil {
			return nil
		}

		if current.next.Value == el {
			next := current.next
			current.next = &Node{Value: v, next: next}
			l.len++
			break
		}

		current = current.next
	}

	return nil
}

func (l *LinkedList) InsertAfter(v, el interface{}) error {
	if l.size == l.len {
		return errors.New("list is full")
	}

	if l.len == 0 {
		return errors.New("list is empty")
	}

	current := l.head
	for current != nil {
		if current.Value == el {
			next := current.next
			current.next = &Node{Value: v, next: next}
			l.len++
			break
		}

		current = current.next
	}

	return nil
}

func (l *LinkedList) Remove(v interface{}) {
	if l.len == 0 {
		return
	}

	if l.head.Value == v {
		l.head = l.head.next
		l.len--
		return
	}

	current := l.head

	for current.next != nil {
		if current.next.Value == v {
			current.next = current.next.next
			l.len--
			return
		}

		current = current.next
	}
}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Println(current.Value)
		current = current.next
	}
}

func (l *LinkedList) Get(v interface{}) *Node {
	current := l.head

	for current != nil {
		if current.Value == v {
			return current
		}

		current = current.next
	}

	return nil
}
