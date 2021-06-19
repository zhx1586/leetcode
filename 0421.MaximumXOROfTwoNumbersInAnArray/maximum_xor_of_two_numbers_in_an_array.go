package maximumxoroftwonumbersinanarray

import (
	"math"
)

func findMaximumXOR(nums []int) int {
	set := map[int]bool{}
	or := 0
	and := int(math.MaxInt32)
	for _, n := range nums {
		set[n] = true
		or = or | n
		and = and & n
	}
	all := or & (^and)

	for t := all; t >= 0; t-- {
		if t|all == all {
			for _, n := range nums {
				if _, ok := set[n^t]; ok {
					return t
				}
			}
		}
	}

	return 0
}
