package main

import (
	"learn-logics/calculator"
	"learn-logics/util"
	"log"
)

func main() {
	a := util.DecimalToBitAsBool(192)
	b := util.DecimalToBitAsBool(65536)
	res := calculator.Adder(a, b)
	log.Println(util.BoolBitToDecimal(res))
	res2 := calculator.Adder(res, util.DecimalToBitAsBool(1_000_000))
	log.Println(util.BoolBitToDecimal(res2))
}
