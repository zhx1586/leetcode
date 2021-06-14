package arithmeticslices

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	ret, dp := 0, 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp = dp + 1
			ret = ret + dp
		} else {
			dp = 0
		}
	}

	return ret
}
