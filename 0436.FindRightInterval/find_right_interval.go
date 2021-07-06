package findrightinterval

import (
	"sort"
)

type Interval struct {
	start, end, index int
}

type Intervals []Interval

func (h Intervals) Len() int           { return len(h) }
func (h Intervals) Less(a, b int) bool { return h[a].start < h[b].start }
func (h Intervals) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }

func findRightInterval(intervals [][]int) []int {
	data := make([]Interval, len(intervals))
	for i := 0; i < len(intervals); i++ {
		data[i] = Interval{
			start: intervals[i][0],
			end:   intervals[i][1],
			index: i,
		}
	}
	sort.Sort(Intervals(data))

	ret := make([]int, len(intervals))
	for i := 0; i < len(intervals); i++ {
		if t := sort.Search(len(intervals), func(curr int) bool {
			return data[curr].start >= intervals[i][1]
		}); t != len(intervals) {
			ret[i] = data[t].index
		} else {
			ret[i] = -1
		}
	}

	return ret
}
