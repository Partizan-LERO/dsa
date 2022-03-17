package sorting

import "github.com/Partizan-LERO/dsa/internal/ds/heap"

// HeapSort sorts an array in descending order
// For an ascending order, please, change struct to minHeap
func HeapSort(arr []int) []int {
	h := heap.MaxHeap{}
	var res []int

	for _, v := range arr {
		h.Insert(v)
	}

	for i := 0; i < len(arr); i++ {
		res = append(res, h.ExtractMax())
	}

	return res
}

// HeapSortASC sorts an array in ascending order
// For a descending order, please, change struct to maxHeap
func HeapSortASC(arr []int) []int {
	h := heap.MinHeap{}
	var res []int

	for _, v := range arr {
		h.Insert(v)
	}

	for i := 0; i < len(arr); i++ {
		res = append(res, h.ExtractMin())
	}

	return res
}
