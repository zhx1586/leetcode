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

	//fmt.Println(len(nums))
	//fmt.Printf("or:%b, and:%b, all:%b\n", or, and, all)

	for t := all; t >= 0; t-- {
		if t|all == all {
			//fmt.Printf("t:%b\n", t)
			for _, n := range nums {
				if _, ok := set[n^t]; ok {
					//fmt.Printf("(%b,%b)\n", n, n^t)
					return t
				}
			}
		}
	}

	return 0
}
