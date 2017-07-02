package LeetCode

import (
	"testing"
	"fmt"
)

func findDerangement(n int) int {
	var counter []int
	counter = append(counter,0) // counter 0
	counter = append(counter,0) // counter 1
	counter = append(counter, 1) // counter 2
	counter = append(counter,2) // counter 3
	var mod = 1000000007
	if n <= 3 {
		return counter[n]
	}
	for i := 4; i <= n; i++ {
		counter = append(counter,(i-1) * ((counter[i-1]+counter[i-2])%(mod)) % mod)
	}
	return counter[n]
}

func TestFindDerangement(t *testing.T)  {
	fmt.Println(findDerangement(4))
}