package main

import (
	"learn-logics/calculator"
	"learn-logics/util"
	"log"
)

func main() {
	a := util.DecimalToBoolBitArray(192)
	b := util.DecimalToBoolBitArray(65536)
	res := calculator.Adder(a, b)
	log.Println(util.BoolBitArrayToDecimal(res))
	res2 := calculator.Adder(res, util.DecimalToBoolBitArray(1_000_000))
	log.Println(util.BoolBitArrayToDecimal(res2))
}
