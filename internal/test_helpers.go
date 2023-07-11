package internal

import (
	"sort"
	"wordStore/internal/utils"
)

func naiveLeast(frequencies map[string]*utils.Element) uint32 {
	var min uint32 = 99999999

	for _, e := range frequencies {
		if e.Frequency < min {
			min = e.Frequency
		}
	}

	if min == 99999999 {
		return 0
	}

	return min
}

func naiveTopK(frequencies map[string]*utils.Element, k int) []*utils.Element {
	var elems utils.Elements

	for _, e := range frequencies {
		elems = append(elems, e)
	}

	sort.Sort(elems)

	res := make([]*utils.Element, 0, k)

	if len(elems) == 0 {
		return res
	}

	for i := 0; i < k && i < len(elems); i++ {
		res = append(res, elems[i])
	}

	return res
}

func naiveMedian(frequencies map[string]*utils.Element) uint32 {
	var elems utils.Elements

	for _, e := range frequencies {
		elems = append(elems, e)
	}

	sort.Sort(elems)

	s := make(utils.Elements, 0, len(elems))

	for i := len(elems) - 1; i >= 0; i-- {
		s = append(s, elems[i])
	}

	if len(s)%2 != 0 {
		return s[len(s)/2].Frequency
	}

	if len(s) == 0 {
		return 0
	}
	return (s[(len(s)-1)/2].Frequency + s[(len(s)/2)].Frequency) / 2

}

func compareElements(elems1 []*utils.Element, elems2 []*utils.Element) bool {
	if len(elems1) != len(elems2) {
		return false
	}

	for i := 0; i < len(elems1); i++ {
		// order of words is not guaranteed hence checking only frequencies is sufficient
		if elems1[i].Frequency != elems2[i].Frequency {
			return false
		}
	}

	return true
}
