package utils

type Elements []*Element

type Element struct {
	Word      string `json:"word"`
	Frequency uint32 `json:"frequency"`

	// Indexes indicating the presence of this element in each of the median heaps
	// this is a tradeoff between time and space as without those fields we would have to iterate over the heaps, and it can be costly.
	SmallerIdx int `json:"-"`
	LargerIdx  int `json:"-"`
}

func ElementCopy(e *Element) *Element {
	if e == nil {
		return nil
	}

	return &Element{
		Word:      e.Word,
		Frequency: e.Frequency,
	}
}

// Less for sorting in descending order
func (s Elements) Less(i, j int) bool {
	return s[i].Frequency > s[j].Frequency
}

func (s Elements) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Elements) Len() int {
	return len(s)
}
