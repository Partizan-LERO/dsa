package main

import (
	"fmt"
	"github.com/Partizan-LERO/dsa/internal/ds/stack_arr"
)

func main() {
	/*l := linked_list.NewLinkedList(5)
	l.Add(1)
	l.InsertBefore("be", 2)
	l.Add("Test2")
	l.Add("Test4")
	l.InsertAfter("After", "Test2")
	l.InsertBefore("Before", 1)
	l.Add("Test5")
	l.Remove("Test2")
	l.Print() // before 1 2 3 4*/

	/*dll := dll2.NewLinkedList(6)
	dll.Add(1)
	dll.InsertBefore("be", 2)
	dll.Add(2)
	dll.Add(3)
	dll.InsertAfter("After", 2)
	dll.InsertBefore("Before", 1)
	dll.Add(4)
	dll.Remove(2)
	dll.Print() */

	/*s := stack.NewStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())*/

	s := stack_arr.NewStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
