package bst

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	}

	if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	}

	if n.Key > k {
		return n.Left.Search(k)
	}

	return true
}
