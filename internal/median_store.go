package internal

func newMedianStore() *medianStore {
	ms := medianStore{
		smaller: NewMaxHeap(), // TODO: change to minMaxHeap
		larger:  NewMinHeap(),
	}

	return &ms
}

func (s *medianStore) getMedian() Element {
	return Element{}
}
