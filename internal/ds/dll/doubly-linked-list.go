package dll

import (
	"errors"
	"fmt"
)

type DoublyLinkedList struct {
	len  int
	size int
	head *Node
}

type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

func NewLinkedList(size int) DoublyLinkedList {
	return DoublyLinkedList{
		size: size,
	}
}

func (l *DoublyLinkedList) Add(v interface{}) error {
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

		current.next = &Node{Value: v, prev: current}
	}

	l.len++
	return nil
}

func (l *DoublyLinkedList) InsertBefore(v, el interface{}) error {
	if l.size == l.len {
		return errors.New("list is full")
	}

	if l.len == 0 {
		return errors.New("list is empty")
	}

	current := l.head

	if current.Value == el {
		l.head = &Node{Value: v, next: current}
		current.prev = l.head
		l.len++
		return nil
	}

	for current != nil {
		if current.next == nil {
			return nil
		}

		if current.next.Value == el {
			next := current.next
			current.next = &Node{Value: v, next: next, prev: current}
			next.prev = current.next
			l.len++
			break
		}

		current = current.next
	}

	return nil
}

func (l *DoublyLinkedList) InsertAfter(v, el interface{}) error {
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
			if current.next.next != nil {
				current.next.next.prev = current
			}
			current.next = &Node{Value: v, next: next, prev: current}
			l.len++
			break
		}

		current = current.next
	}

	return nil
}

func (l *DoublyLinkedList) Remove(v interface{}) {
	if l.len == 0 {
		return
	}

	if l.head.Value == v {
		l.head = l.head.next
		l.head.prev = nil
		l.len--
		return
	}

	current := l.head

	for current.next != nil {
		if current.next.Value == v {
			current.next.prev = current
			current.next = current.next.next
			l.len--
			return
		}

		current = current.next
	}
}

func (l *DoublyLinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Println(current.Value)
		current = current.next
	}
}

func (l *DoublyLinkedList) Get(v interface{}) *Node {
	current := l.head

	for current != nil {
		if current.Value == v {
			return current
		}

		current = current.next
	}

	return nil
}
