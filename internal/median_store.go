package internal

import (
	"container/heap"
	"sync"
	"wordStore/internal/utils"
)

func newMedianStore() *medianStore {
	ms := medianStore{
		smaller: utils.NewMinMaxHeap("smaller"),
		larger:  utils.NewMinHeap("larger"),
	}

	return &ms
}

// insert updates medianStore with the new frequencies
func (s *medianStore) insert(elements map[string]*utils.Element, wg *sync.WaitGroup) {
	defer wg.Done()

	// fix heaps in-case at least one element's frequency changed
	for _, e := range elements {
		if s.inSmallHeapSet(e) {
			utils.Fix(s.smaller, e.SmallerIdx)
			continue
		}

		if s.inLargeHeapSet(e) {
			heap.Fix(s.larger, e.LargerIdx)
		}
	}

	for _, e := range elements {
		if !s.inSmallHeapSet(e) && !s.inLargeHeapSet(e) {
			utils.Push(s.smaller, e)
		}

		// make sure every element in small is <= every element in large
		for s.smaller.Len() > 0 && s.larger.Len() > 0 &&
			s.smaller.PeekMax().Frequency > s.larger.PeekMin().Frequency {

			// remove max from smaller heap
			maxFromSmaller := utils.PopMax(s.smaller).(*utils.Element)

			// add max to larger heap
			heap.Push(s.larger, maxFromSmaller)
		}

		for s.smaller.Len() > s.larger.Len()+1 {
			// remove max from smaller heap
			maxFromSmaller := utils.PopMax(s.smaller).(*utils.Element)

			// add max to larger heap
			heap.Push(s.larger, maxFromSmaller)
		}

		for s.larger.Len() > s.smaller.Len()+1 {
			// remove min from larger heap
			minFromLarger := heap.Pop(s.larger).(*utils.Element)

			// add min to smaller heap
			utils.Push(s.smaller, minFromLarger)
		}
	}

}

// getMedian returns the median
func (s *medianStore) getMedian() uint32 {
	if s.smaller.Len() > s.larger.Len() {
		return s.smaller.PeekMax().Frequency
	}

	if s.larger.Len() > s.smaller.Len() {
		return s.larger.PeekMin().Frequency
	}

	if s.larger.Len() == 0 && s.smaller.Len() == 0 {
		return 0
	}

	return (s.smaller.PeekMax().Frequency + s.larger.PeekMin().Frequency) / 2
}

// getLeast returns the least frequent word by leveraging the MinMax heap data structure.
// the max aspect of it is responsible for the median and the min for the least
func (s *medianStore) getLeast() uint32 {
	var min *utils.Element

	if s.smaller.Len() > 0 {
		min = utils.Pop(s.smaller).(*utils.Element) // removes min
		utils.Push(s.smaller, min)
	}

	if min == nil {
		return 0
	}

	return min.Frequency
}

func (s *medianStore) inSmallHeapSet(e *utils.Element) bool {
	return e.SmallerIdx >= 0
}

func (s *medianStore) inLargeHeapSet(e *utils.Element) bool {
	return e.LargerIdx >= 0
}
