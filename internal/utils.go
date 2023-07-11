package internal

type Heap struct {
	Elements []*Element
	less     func(elements []*Element, i, j int) bool
}

func (h Heap) Len() int { return len(h.Elements) }

func (h Heap) Swap(i, j int) { h.Elements[i], h.Elements[j] = h.Elements[j], h.Elements[i] }

func (h Heap) Less(i, j int) bool { return h.less(h.Elements, i, j) }

func (h *Heap) Push(x any) {
	h.Elements = append(h.Elements, x.(*Element))
}

func (h *Heap) Pop() any {
	old := h.Elements
	n := len(old)
	x := old[n-1]
	h.Elements = old[0 : n-1]

	return x
}

func (h *Heap) Peek() any {
	return h.Elements[h.Len()-1]
}

func NewMinHeap() *Heap {
	return &Heap{
		Elements: []*Element{},
		less: func(elements []*Element, i, j int) bool {
			return elements[i].Frequency < elements[j].Frequency
		},
	}
}

func NewMaxHeap() *Heap {
	return &Heap{
		Elements: []*Element{},
		less: func(elements []*Element, i, j int) bool {
			return elements[i].Frequency > elements[j].Frequency
		},
	}
}

func elementCopy(e *Element) Element {
	return Element{
		Word:      e.Word,
		Frequency: e.Frequency,
	}
}
