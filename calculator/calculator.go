package calculator

import (
	"learn-logics/adder"
	"learn-logics/util"
)

func Adder(a, b []bool) []bool {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	if len(a) > len(b) {
		diff := len(a) - len(b)
		padding := make([]bool, diff)
		b = append(padding, b...)
	} else if len(b) > len(a) {
		diff := len(b) - len(a)
		padding := make([]bool, diff)
		a = append(padding, a...)
	}

	c, s := adder.HalfAdder(a[maxLen-1], b[maxLen-1])
	ret := []bool{s}
	for i := maxLen - 2; i >= 0; i-- {
		c, s = adder.FullAdder(a[i], b[i], c)
		ret = append(ret, s)
	}
	if c {
		ret = append(ret, c)
	}
	return util.Reverse(ret)
}
