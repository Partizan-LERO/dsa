package heap

type MaxHeap struct {
	items []int
}

func (h *MaxHeap) Insert(n int) {
	h.items = append(h.items, n)

	h.heapifyUp(len(h.items) - 1)
}

func (h *MaxHeap) ExtractMax() int {
	l := len(h.items) - 1

	if l == -1 {
		return -1
	}

	max := h.items[0]

	h.swap(l, 0)
	h.items = h.items[:l]

	if l > 0 {
		h.heapifyDown(0)
	}

	return max
}

func (h *MaxHeap) heapifyDown(index int) {
	maxElIndex := h.left(index)
	if h.left(index) > len(h.items)-1 {
		return
	}

	if h.right(index) > len(h.items)-1 {
		maxElIndex = h.left(index)
	} else if h.items[h.left(index)] < h.items[h.right(index)] {
		maxElIndex = h.right(index)
	}

	if h.items[index] < h.items[maxElIndex] {
		h.swap(index, maxElIndex)
		h.heapifyDown(maxElIndex)
	}
}

func (h *MaxHeap) heapifyUp(index int) {
	if index == 0 {
		return
	}

	for h.items[index] > h.items[h.parent(index)] {
		h.swap(index, h.parent(index))
		h.heapifyUp(h.parent(index))
	}
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.items[i1], h.items[i2] = h.items[i2], h.items[i1]
}

func (h *MaxHeap) left(i int) int {
	return i*2 + 1
}

func (h *MaxHeap) right(i int) int {
	return i*2 + 2
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}
