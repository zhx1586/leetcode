package splitarraylargestsum

func splitArray(nums []int, m int) int {
	cache := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		cache[i] = make([]int, len(nums))
	}
	for x := 0; x < len(nums); x++ {
		cache[x][x] = nums[x]
		for y := x + 1; y < len(nums); y++ {
			cache[x][y] = cache[x][y-1] + nums[y]
		}
	}

	ret := cache[0][len(nums)-1]
	splitArrayCore(0, m, &[]int{0}, &cache, &ret)

	return ret
}

func splitArrayCore(curr int, m int, path *[]int, cache *[][]int, ret *int) {
	if m == 1 {
		*path = append(*path, len(*cache))
		largest := 0
		for i := 1; i < len(*path); i++ {
			sum := (*cache)[(*path)[i-1]][(*path)[i]-1]
			if largest < sum {
				largest = sum
			}
		}
		if *ret > largest {
			*ret = largest
		}
		*path = (*path)[0 : len(*path)-1]
		return
	}
	for next := curr + 1; next < len(*cache); next++ {
		*path = append(*path, next)
		splitArrayCore(next, m-1, path, cache, ret)
		*path = (*path)[0 : len(*path)-1]
	}
}
