package internal

import (
	"container/heap"
	"sort"
	"sync"
	"wordStore/internal/utils"
)

func newTopKStore(k int) *topKStore {
	tks := topKStore{
		k:    k,
		heap: utils.NewMinHeap("topK"),
	}

	return &tks
}

// insert updates topKStore with the new frequencies
func (s *topKStore) insert(elements map[string]*utils.Element, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, elem := range elements {
		var found bool

		for i, e := range s.heap.Elements { // looping takes O(k) ~ O(1) when k << store_size
			if elem.Word == e.Word { // word already present in heap, heapify
				heap.Fix(s.heap, i)
				found = true
				break
			}
		}

		if found {
			continue
		}

		if s.heap.Len() < s.k {
			heap.Push(s.heap, elem)
			continue
		}

		min := s.heap.PeekMin()

		if min.Frequency < elem.Frequency {
			heap.Pop(s.heap)
			heap.Push(s.heap, elem)
		}
	}
}

// getTopK returns at most k most frequent words
func (s *topKStore) getTopK() []*utils.Element {
	topCopy := make(utils.Elements, s.heap.Len())

	for i, e := range s.heap.Elements {
		topCopy[i] = utils.ElementCopy(e)
	}

	// sort in desc order
	sort.Sort(topCopy)

	return topCopy
}
