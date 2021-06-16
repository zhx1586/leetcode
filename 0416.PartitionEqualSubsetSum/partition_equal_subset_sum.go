package partitionequalsubsetsum

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum = sum + n
	}
	if sum%2 == 1 {
		return false
	}

	dp := [][]bool{}
	for i := 0; i < len(nums)+1; i++ {
		dp = append(dp, make([]bool, sum+1))
	}
	dp[0][0] = true

	for i := 0; i < len(nums); i++ {
		for j := 1; j <= sum; j++ {
			if j >= nums[i] {
				dp[i+1][j] = dp[i][j-nums[i]] || dp[i][j]
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}

	//fmt.Println(nums)
	//for _, r := range dp {
	//fmt.Println(r)
	//}

	return dp[len(nums)][sum/2]
}
