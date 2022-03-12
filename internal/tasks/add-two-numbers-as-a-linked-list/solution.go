package add_two_numbers_as_a_linked_list

import (
	"strconv"
)

type Node struct {
	Value int
	Next  *Node
}

func getNumber(l Node) int {
	current := l
	v := ""
	for true {
		v += strconv.Itoa(current.Value)
		if current.Next == nil {
			break
		}
		current = *current.Next
	}

	res, _ := strconv.Atoi(v)
	return res
}

func solution(l1, l2 Node) *Node {
	a := getNumber(l1)
	b := getNumber(l2)

	sum := strconv.Itoa(a + b)
	res := &Node{}

	current := res

	for _, l := range sum {
		v, _ := strconv.Atoi(string(l))
		current.Next = &Node{Value: v}
		current = current.Next
	}

	return res.Next
}

func Solution() {
	l1 := Node{Value: 5, Next: &Node{Value: 3, Next: &Node{Value: 2}}}
	l2 := Node{Value: 4, Next: &Node{Value: 5, Next: &Node{Value: 6}}}

	l3 := solution(l1, l2)
	current := l3
	for true {
		if current.Next == nil {
			break
		}
		current = current.Next
	}
}
