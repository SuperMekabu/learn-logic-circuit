package basicCircuit

func NAND(a, b bool) bool {
	return !(a && b) || (a && !b) || (!a && b)
}

func NOT(a bool) bool {
	return NAND(a, a)
}

func AND(a, b bool) bool {
	return NOT(NAND(a, b))
}

func NOR(a, b bool) bool {
	return NOT(NAND(NOT(a), NOT(b)))
}

func OR(a, b bool) bool {
	return NAND(NOT(a), NOT(b))
}

func XOR(a, b bool) bool {
	return AND(NAND(a, b), OR(a, b))
}

func BigOR(a, b, c bool) bool {
	return OR(OR(a, b), c)
}

func BigAND(a, b, c bool) bool {
	return AND(AND(a, b), c)
}

func XNOR(a, b bool) bool {
	return NOT(XOR(a, b))
}
