package internal

import (
	"container/heap"
	"sort"
)

func newTopKStore(k int) *topKStore {
	tks := topKStore{
		k:    k,
		heap: NewMinHeap(),
	}

	return &tks
}

func (s *topKStore) insert(elem *Element) {
	for i, e := range s.heap.Elements { // looping takes O(k) ~ O(1) when k << store_size
		if elem.Word == e.Word { // word already present in heap, heapify
			heap.Fix(s.heap, i)
			return
		}
	}

	if s.heap.Len() < s.k {
		heap.Push(s.heap, elem)
		return
	}

	min := s.heap.Peek().(*Element)

	if min.Frequency < elem.Frequency {
		heap.Pop(s.heap)
		heap.Push(s.heap, elem)
	}

}

func (s *topKStore) getTopK() []Element {
	topCopy := make([]Element, s.heap.Len())

	for i, e := range s.heap.Elements {
		topCopy[i] = elementCopy(e)
	}

	// sort in desc order
	sort.Slice(topCopy, func(i, j int) bool {
		return topCopy[i].Frequency > topCopy[i].Frequency
	})

	return topCopy
}
