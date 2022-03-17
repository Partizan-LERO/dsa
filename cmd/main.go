package main

import (
	"fmt"
	"github.com/Partizan-LERO/dsa/internal/ds/bst"
	"github.com/Partizan-LERO/dsa/internal/ds/graphs"
	"github.com/Partizan-LERO/dsa/internal/ds/hashmap"
	"github.com/Partizan-LERO/dsa/internal/ds/heap"
	"github.com/Partizan-LERO/dsa/internal/ds/stack_arr"
	"github.com/Partizan-LERO/dsa/internal/ds/trie"
	"github.com/Partizan-LERO/dsa/internal/sorting"
	add_two_numbers_as_a_linked_list "github.com/Partizan-LERO/dsa/internal/tasks/add-two-numbers-as-a-linked-list"
	first_and_last_indices_of_an_element_in_a_sorted_array "github.com/Partizan-LERO/dsa/internal/tasks/first-and-last-indices-of-an-element-in-a-sorted-array"
	permutations2 "github.com/Partizan-LERO/dsa/internal/tasks/permutations"
	ransom_note "github.com/Partizan-LERO/dsa/internal/tasks/ransom-note"
	sorting_with_three_unique_numbers "github.com/Partizan-LERO/dsa/internal/tasks/sorting-with-three-unique-numbers"
	two_sum "github.com/Partizan-LERO/dsa/internal/tasks/two-sum"
)

func linkedList() {
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
}

func dll() {
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
}

func stack() {
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
}

func stackArr() {
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

func ransomNote() {
	ransom_note.Solution()
}

func addTwoNumberAsLinkedList() {
	add_two_numbers_as_a_linked_list.Solution()
}

func twoSum() {
	two_sum.Solution()
}

func findFirstAndLast() {
	first_and_last_indices_of_an_element_in_a_sorted_array.Solution()
}

func permutations() {
	permutations2.Solution()
}

func sortingWithThreeUnique() {
	sorting_with_three_unique_numbers.Solution()
}

func testGraph() {
	g := graphs.Graph{}
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.Print()
}

func testGraphM() {
	g := graphs.NewGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.Print()
}

func testBST() {
	tree := bst.Node{Key: 100}
	tree.Insert(101)
	tree.Insert(102)
	tree.Insert(95)
	tree.PrintPreorder()

	fmt.Println(tree.Search(100))
	fmt.Println(tree.Search(101))
	fmt.Println(tree.Search(94))
	fmt.Println(tree.Search(95))
}

func testTrie() {
	t := trie.InitTrie()
	t.Insert("test")
	fmt.Println(t.Search("test"))
	fmt.Println(t.Search("tes"))
	t.Insert("tes")

	fmt.Println(t.Search("tes"))

	t.Insert("march")
	t.Insert("sword")
	fmt.Println(t.Search("march"))
	fmt.Println(t.Search("sword"))

}

func testSelectionSort() {
	arr := []int{64, 12, 25, 12, 22, 11}
	fmt.Println(sorting.SelectionSort(arr))
}

func testBubbleSort() {
	arr := []int{64, 12, 25, 12, 22, 11}
	fmt.Println(sorting.BubbleSort(arr))
}

func testInsertionSort() {
	arr := []int{64, 12, 25, 12, 22, 11}
	fmt.Println(sorting.InsertionSort(arr))
}

func testHashMap() {
	h := hashmap.InitHashMap(2)

	h.Set("name", "Sergei")
	h.Set("same", "YES")
	fmt.Println(h.Get("name"))
	fmt.Println(h.Get("same"))
	h.Delete("name")
	fmt.Println("name", h.Get("name"))
	h.Delete("name")

	fmt.Println(h.Get("age"))
	h.Set("age", "30")
	fmt.Println(h.Get("age"))

	h.Set("same", "False")
	fmt.Println(h.Get("same"))
}

func testHeap() {
	maxHeap := heap.MaxHeap{}

	elements := []int{10, 20, 30, 15, 21, 22, 25, 35}
	for _, v := range elements {
		maxHeap.Insert(v)
		fmt.Println(maxHeap)
	}

	for _, _ = range elements {
		fmt.Println(maxHeap.ExtractMax())
		fmt.Println(maxHeap)
	}
}

func testHeapSort() {
	elements := []int{10, 20, 30, 15, 21, 22, 25, 35}
	sorted := sorting.HeapSort(elements)
	fmt.Println(sorted)

	sorted = sorting.HeapSortASC(elements)
	fmt.Println(sorted)
}

func testMinHeap() {
	minHeap := heap.MinHeap{}

	elements := []int{10, 20, 30, 15, 21, 22, 25, 35}
	for _, v := range elements {
		minHeap.Insert(v)
		fmt.Println(minHeap)
	}

	for _, _ = range elements {
		fmt.Println(minHeap.ExtractMin())
		fmt.Println(minHeap)
	}
}

func testMergeSort() {
	elements := []int{10, 20, 30, 15, 21, 22, 25, 35}
	fmt.Println(sorting.MergeSort(elements))
}

func main() {
	testMergeSort()
}
