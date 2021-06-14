package splitarraylargestsum

func splitArrayV3(nums []int, m int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		if i == 0 {
			dp[i][0] = nums[i]
		} else {
			dp[i][0] = dp[i-1][0] + nums[i]
		}
	}
	//fmt.Printf("nums:%+v\n", nums)
	//fmt.Printf("dp:%+v\n", dp)

	for x := 2; x <= m; x++ {
		for y := x - 1; y < n; y++ {
			//fmt.Printf("=== (x: %d, y: %d) ===\n", x, y)
			sum := nums[y]
			dp[y][x-1] = max(dp[y-1][x-2], sum)
			//fmt.Printf("sum: %d, dp[%d][%d]: %d, max(%d, %d)\n", sum, y, x-1, dp[y][x-1], dp[y-1][x-2], sum)
			for z := y - 1; z >= x-1; z-- {
				sum = sum + nums[z]
				dp[y][x-1] = min(dp[y][x-1], max(dp[z-1][x-2], sum))
				if sum > dp[y][x-1] {
					break
				}
				//fmt.Printf("sum: %d, dp[%d][%d]: %d, max(%d, %d)\n", sum, y, x-1, dp[y][x-1], dp[z-1][x-2], sum)
			}
		}
	}

	return dp[n-1][m-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
