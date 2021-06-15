package partitionequalsubsetsum

import "sort"

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum = sum + n
	}
	if sum%2 == 1 {
		return false
	}

	sort.Sort(sort.IntSlice(nums))

	return canPartitionCore(nums, 0, sum/2)
}

func canPartitionCore(nums sort.IntSlice, pos int, target int) bool {
	if pos == len(nums) {
		return false
	}
	if nums[pos] > target {
		return false
	} else if nums[pos] < target {
		return canPartitionCore(nums, pos+1, target) ||
			canPartitionCore(nums, pos+1, target-nums[pos])
	} else {
		return true
	}
}
