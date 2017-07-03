package LeetCode

import "testing"

func findMin(nums []int) int {
	// note: 可用二分法！！
	if len(nums) == 1 {
		return nums[0]
	}
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}
	return nums[0]
}

func TestFindMin(t *testing.T) {
	if findMin([]int{1}) != 1 {
		t.Error("test1 error")
	}
	if findMin([]int{4,5,6,7,0,1,2,3}) != 0 {
		t.Error("wrong")
	}
}