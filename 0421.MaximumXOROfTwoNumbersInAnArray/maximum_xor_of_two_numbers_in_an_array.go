package maximumxoroftwonumbersinanarray

import "math"

func findMaximumXOR(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	dp := nums[0]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]^nums[j] > dp {
				dp = nums[i] ^ nums[j]
				if dp == int(math.MaxInt32) {
					return dp
				}
			}
		}
	}

	return dp
}
