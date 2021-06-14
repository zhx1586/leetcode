package thirdmaximumnumber

import "container/heap"

type intHeap []int

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func (h intHeap) Has(x int) bool {
	for _, elem := range h {
		if elem == x {
			return true
		}
	}
	return false
}

func thirdMax(nums []int) int {
	maximum := make(intHeap, 0)

	for _, n := range nums {
		if maximum.Has(n) {
			continue
		}
		if len(maximum) < 3 {
			heap.Push(&maximum, n)
		} else {
			if n > maximum[len(maximum)-1] {
				heap.Pop(&maximum)
				heap.Push(&maximum, n)
			}
		}
	}

	if len(maximum) == 2 {
		heap.Pop(&maximum)
	}
	return heap.Pop(&maximum).(int)
}
