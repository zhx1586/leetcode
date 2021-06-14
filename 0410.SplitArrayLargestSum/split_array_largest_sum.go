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
	splitArrayCore(0, m, 0, &cache, &ret)

	return ret
}

func splitArrayCore(curr int, m int, largest int, cache *[][]int, ret *int) {
	if m == 1 {
		sum := (*cache)[curr][len(*cache)-1]
		if largest < sum {
			largest = sum
		}
		if *ret > largest {
			*ret = largest
		}
		return
	}
	for next := curr + 1; next < len(*cache); next++ {
		sum := (*cache)[curr][next-1]
		if sum > *ret {
			return
		}
		if largest < sum {
			splitArrayCore(next, m-1, sum, cache, ret)
		} else {
			splitArrayCore(next, m-1, largest, cache, ret)
		}
	}
}
