package nonoverlappingintervals

import "sort"

type Intervals [][]int

func (t Intervals) Len() int { return len(t) }
func (t Intervals) Less(i, j int) bool {
	if t[i][0] < t[j][0] {
		return true
	} else if t[i][0] > t[j][0] {
		return false
	} else {
		return t[i][1] < t[j][1]
	}
}
func (t Intervals) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Sort(Intervals(intervals))

	count, last := 0, 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[last][1] {
			count++
			if intervals[i][1] < intervals[last][1] {
				last = i
			}
		} else {
			last = i
		}
	}

	return count
}
