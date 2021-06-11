package queuereconstructionbyheight

import (
	"fmt"
	"sort"
)

type People [][]int

func (h People) Len() int { return len(h) }
func (h People) Less(i, j int) bool {
	if h[i][0] < h[j][0] {
		return true
	} else if h[i][0] > h[j][0] {
		return false
	} else {
		return h[i][1] < h[j][1]
	}
}
func (h People) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func reconstructQueue(people [][]int) [][]int {
	sort.Sort(People(people))
	fmt.Println(people)
	ret := make([][]int, len(people))
	for _, p := range people {
		k := p[1]
		pos := 0
		for {
			if k == 0 && len(ret[pos]) == 0 {
				ret[pos] = p
				break
			} else if len(ret[pos]) > 0 && ret[pos][0] < p[0] {
				pos++
			} else {
				pos++
				k--
			}
		}
	}
	return ret
}
