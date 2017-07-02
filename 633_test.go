package LeetCode

import (
	"math"
	"testing"
)

func judgeSquareSum(c int) bool {

	high := int(math.Sqrt(float64(c)))
	low := high / 2

	for i := high; i >= low; i-- {
		left := c - i * i
		j := int(math.Sqrt(float64(left)))
		if j*j == left {
			return true
		}
	}
	return false
}

func TestJudgeSquareSum(t *testing.T)  {
	if judgeSquareSum(3) {
		t.Error("test1 false")
	}
	if !judgeSquareSum(5) {
		t.Error("test2 false")
	}
}