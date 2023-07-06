package adder

import (
	"learn-logics/basicCircuit"
)

// HalfAdder
/*
returns CAR, SUM
*/
func HalfAdder(a, b bool) (c bool, s bool) {
	return basicCircuit.AND(a, b), basicCircuit.XOR(a, b)
}

// FullAdder
/*
returns CAR, SUM
*/
func FullAdder(a, b, x bool) (c bool, s bool) {
	// 1段めの半加算器
	f1, f2 := HalfAdder(a, b)
	// 2段めの半加算器
	// 1段めのSUMとc
	s1, s2 := HalfAdder(f2, x)
	// 1段め、2段めのCARと2段めのSUM
	return basicCircuit.OR(f1, s1), s2
}
