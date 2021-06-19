package maximumxoroftwonumbersinanarray

func findMaximumXOR(nums []int) int {
	set := map[int]bool{}
	all := 0
	for _, n := range nums {
		set[n] = true
		all = all | n
	}

	//fmt.Println(len(nums))

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
