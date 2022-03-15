package hashmap

type HashMap struct {
	items []*Buckets
	size  int
}

type Buckets struct {
	head *Item
	tail *Item
}

type Item struct {
	Key   string
	Value string
	Next  *Item
}

func InitHashMap(size int) HashMap {
	return HashMap{
		items: make([]*Buckets, size, size),
		size:  size,
	}
}

func (h *HashMap) Hash(key string) int {
	hash := 7

	for i := 0; i < len(key); i++ {
		hash = hash*31 + i
	}

	if hash > h.size {
		hash = hash % h.size
	}

	return hash
}

func (h *HashMap) Set(key string, value string) {
	index := h.Hash(key)

	item := &Item{
		Key:   key,
		Value: value,
	}

	if h.items[index] != nil {
		h.items[index].tail.Next = item

		h.items[index].tail = h.items[index].tail.Next
	} else {
		h.items = append(h.items[:index+1], h.items[index:]...)

		h.items[index] = &Buckets{
			head: item,
			tail: item,
		}
	}
}

func (b *Buckets) delete(key string) {
	current := b.head

	//head
	if current.Key == key {
		b.head = b.head.Next
	}

	for current != nil {
		if current.Key == key {
			current.Key = current.Next.Key
			current.Value = current.Next.Value
			current.Next = current.Next.Next
			return
		}

		current = current.Next
	}

}

func (h *HashMap) Delete(key string) {
	index := h.Hash(key)

	if h.items[index] != nil {
		h.items[index].delete(key)
	}
}

func (h *HashMap) Get(key string) string {
	index := h.Hash(key)
	item := h.items[index]
	if item == nil {
		return ""
	}

	current := item.head

	for current != nil {
		if current.Key == key {
			return current.Value
		}

		current = current.Next
	}

	return ""
}
