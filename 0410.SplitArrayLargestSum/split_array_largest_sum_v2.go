package splitarraylargestsum

func splitArrayV2(nums []int, m int) int {
	n := len(nums)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := i; j < n; j++ {
			dp[i][j] = make([]int, min(m, j-i+1))
		}
		dp[i][i][0] = nums[i]
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dp[i][j][0] = dp[i][j-1][0] + nums[j]
		}
	}
	for k := 2; k <= m; k++ {
		for x := 0; x < n; x++ {
			for y := x + k - 1; y < n; y++ {
				ans := dp[x][y][0]
				for j := x; j < y && y-j >= k-1; j++ {
					sum := max(dp[x][j][0], dp[j+1][y][k-2])
					if ans > sum {
						ans = sum
					}
				}
				dp[x][y][k-1] = ans
			}
		}
	}
	return dp[0][n-1][m-1]
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
