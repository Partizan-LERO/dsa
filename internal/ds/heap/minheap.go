package heap

type MinHeap struct {
	items []int
}

func (h *MinHeap) Insert(n int) {
	h.items = append(h.items, n)

	h.heapifyUp(len(h.items) - 1)
}

func (h *MinHeap) ExtractMin() int {
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

func (h *MinHeap) heapifyDown(index int) {
	minElIndex := h.left(index)
	if h.left(index) > len(h.items)-1 {
		return
	}

	if h.right(index) > len(h.items)-1 {
		minElIndex = h.left(index)
	} else if h.items[h.left(index)] > h.items[h.right(index)] {
		minElIndex = h.right(index)
	}

	if h.items[index] > h.items[minElIndex] {
		h.swap(index, minElIndex)
		h.heapifyDown(minElIndex)
	}
}

func (h *MinHeap) heapifyUp(index int) {
	if index == 0 {
		return
	}

	for h.items[index] < h.items[h.parent(index)] {
		h.swap(index, h.parent(index))
		h.heapifyUp(h.parent(index))
	}
}

func (h *MinHeap) swap(i1, i2 int) {
	h.items[i1], h.items[i2] = h.items[i2], h.items[i1]
}

func (h *MinHeap) left(i int) int {
	return i*2 + 1
}

func (h *MinHeap) right(i int) int {
	return i*2 + 2
}

func (h *MinHeap) parent(i int) int {
	return (i - 1) / 2
}
