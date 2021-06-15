package partitionequalsubsetsum

import (
	"sort"
)

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
	if target == 0 {
		return true
	} else if target < 0 {
		return false
	}
	if pos >= len(nums) {
		return false
	}
	count := 1
	for i := pos + 1; i < len(nums); i++ {
		if nums[i] > nums[pos] {
			break
		}
		count++
	}
	for i := 1; i <= count; i++ {
		if canPartitionCore(nums, pos+count, target-i*nums[pos]) {
			return true
		}
	}
	return false
}
