package LeetCode

import (
	"fmt"
	"testing"
)

func addBinary(a string, b string) string {
	aindex := len(a) - 1
	bindex := len(b) - 1

	var res string
	var c int = 0
	for ; aindex >= 0 && bindex >= 0; {
		anum := int(a[aindex] - '0')
		bnum := int(b[bindex] - '0')
		juge := anum + bnum + c
		c = juge / 2
		res = fmt.Sprintf("%d%s",juge%2,res)
		bindex--
		aindex--
	}

	for aindex >= 0 {
		anum := int(a[aindex] - '0')
		res = fmt.Sprintf("%d%s",(anum+c)%2,res)
		c = (anum + c) / 2
		aindex--
	}

	for bindex >= 0 {
		bnum := int(b[bindex] - '0')
		res = fmt.Sprintf("%d%s",(bnum+c)%2,res)
		c = (bnum + c) / 2
		bindex--
	}
	if c == 1 {
		res = fmt.Sprintf("1%s",res)
	}
	return res
}

func TestAddBinary(t *testing.T) {
	a := "101111"
	b := "10"
	c := "110001"
	if c != addBinary(a, b) {
		t.Errorf("Expert %s, Get %s",c, addBinary(a,b))
	}
}