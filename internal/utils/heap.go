package utils

// min-heap implementation

type Heap struct {
	Elements []*Element
	less     func(elements []*Element, i, j int) bool
	id       string
}

func (h Heap) Len() int { return len(h.Elements) }

func (h Heap) Swap(i, j int) {
	h.Elements[i], h.Elements[j] = h.Elements[j], h.Elements[i]

	if h.id == "larger" {
		h.Elements[i].LargerIdx, h.Elements[j].LargerIdx = h.Elements[j].LargerIdx, h.Elements[i].LargerIdx
	}

	if h.id == "smaller" {
		h.Elements[i].SmallerIdx, h.Elements[j].SmallerIdx = h.Elements[j].SmallerIdx, h.Elements[i].SmallerIdx
	}
}

func (h Heap) Less(i, j int) bool { return h.less(h.Elements, i, j) }

func (h *Heap) Push(x any) {
	e := x.(*Element)
	if h.id == "larger" {
		e.LargerIdx = h.Len()
	}

	if h.id == "smaller" {
		e.SmallerIdx = h.Len()
	}

	h.Elements = append(h.Elements, e)
}

func (h *Heap) Pop() any {
	old := h.Elements
	n := len(old)
	x := old[n-1]

	if h.id == "larger" {
		x.LargerIdx = -1
	}

	if h.id == "smaller" {
		x.SmallerIdx = -1
	}

	h.Elements = old[0 : n-1]

	return x
}

func (h *Heap) PeekMin() *Element {
	return h.Elements[0]
}

func NewMinHeap(t string) *Heap {
	return &Heap{
		Elements: []*Element{},
		less: func(elements []*Element, i, j int) bool {
			return elements[i].Frequency < elements[j].Frequency
		},
		id: t,
	}
}
