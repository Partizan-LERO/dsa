package trie

const AbcLength = 26

type Node struct {
	Children [AbcLength]*Node
	IsEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

func (t *Trie) Insert(w string) {
	wLen := len(w)
	currentNode := t.root
	for i := 0; i < wLen; i++ {
		charIndex := w[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			currentNode.Children[charIndex] = &Node{}
		}

		currentNode = currentNode.Children[charIndex]
	}

	currentNode.IsEnd = true
}

func (t *Trie) Search(w string) bool {
	wLen := len(w)
	currentNode := t.root
	for i := 0; i < wLen; i++ {
		charIndex := w[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			return false
		}

		currentNode = currentNode.Children[charIndex]
	}

	if currentNode.IsEnd == true {
		return true
	}

	return false
}
