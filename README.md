# Lean the logic circuit on Golang

## Basic circuit
- NOT
- AND
- NAND
- OR
- NOR
- XOR
- XNOR
- BigOR
- BigAND

## Adders
- HalfAdder
  - half adder
- FullAdder
  - full adder

## Calculator
- Adder
  - sum two provided bool slice that as bit array
  - decimals and negative number is not supported

## Utils
- DecimalToBoolAsBitArray
  - convert decimal to bool slice that as bit array
- BoolBitArrayToDecimal
  - convert slice of bool that indicates bit to decimal number
- Reverse
  - reverse order of slice

## Usage
- convert two int32 numbers to two bool slice by DecimalToBoolBitArray
- give those to calculator.Adder
- convert its response by util.BoolBitArrayToDecimal