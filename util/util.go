package util

import (
	"fmt"
	"log"
	"strconv"
)

func DecimalToBitAsBool(src int32) []bool {
	bin := fmt.Sprintf("%b", src)
	var ret []bool
	for _, b := range bin {
		c := fmt.Sprintf("%c", b)
		appender := true
		if c == "0" {
			appender = false
		}
		ret = append(ret, appender)
	}
	return ret
}

func BoolBitToDecimal(src []bool) int64 {
	tmpStr := ""
	for _, c := range src {
		b := 0
		if c {
			b = 1
		}
		tmpStr = fmt.Sprintf("%s%d", tmpStr, b)
	}
	if ret, err := strconv.ParseInt(tmpStr, 2, 0); err != nil {
		log.Fatalf(err.Error())
		return 0
	} else {
		return ret
	}
}

func Reverse[T any](src []T) []T {
	l := src
	for i := 0; i < len(l)/2; i++ {
		l[i], l[len(l)-i-1] = l[len(l)-i-1], l[i]
	}
	return l
}
