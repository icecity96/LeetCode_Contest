package LeetCode

import (
	"sort"
	"testing"
)

func maximumProduct(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	sort.Ints(nums)
	max1 := nums[0] * nums[1] * nums[len(nums)-1]
	max2 := nums[len(nums)-2] * nums[len(nums)-3] * nums[len(nums)-1]
	if max1 > max2 {
		return max1
	}
	return max2
}

func Test628(t *testing.T) {
}
