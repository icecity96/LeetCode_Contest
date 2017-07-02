package LeetCode

import (
	"fmt"
	"testing"
)

type element struct {
	num int
	k int
}

func smallestRange(nums [][]int) []int {
	var k = len(nums)
	var flag = true
	var single []element
	var resmin int
	var resmax int
	for flag  {
		flag = false
		ele := element{1000000,1}
		//=====================================
		// find min 这儿应该用堆！！！否则timeout
		for i := 0; i < k; i++ {
			// list not nil
			if len(nums[i]) != 0 {
				flag = true
				if nums[i][0] < ele.num {
					ele.num = nums[i][0]
					ele.k = i
				}
			}
		}
		//======================================
		if flag == false {
			break
		}
		single = append(single,ele)
		if len(nums[ele.k]) == 1 {
			nums[ele.k] = []int{}
		} else {
			nums[ele.k] = nums[ele.k][1:] // delete first ele
		}
	}

	var pair = make(map[int]int)
	var distance int = 10000000
	for index,v := range single {
		pair[v.k] = index

		if len(pair) == k {
			// map find min index
			min := index + 1
			for _, val := range pair {
				if val < min {
					min = val
				}
			}
			// compare diff
			newdisrance := single[index].num - single[min].num
			if distance > newdisrance {
				distance = single[index].num - single[min].num
				resmax = index
				resmin = min
			}
		}
	}
	return []int{single[resmin].num,single[resmax].num}
}

func TestSmallestRange(t *testing.T) {
	nums := [][]int{
		{4,10,15,24,26},
		{0,9,12,20},
		{5,18,22,30},
	}
	fmt.Println(smallestRange(nums))
}
