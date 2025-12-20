package main

type Operator string

type OperatorList struct {
	Operands []int
	Operator Operator
}

func (ol OperatorList) Evaluate() int {
	result := ol.Operands[0]
	for _, operand := range ol.Operands[1:] {
		switch ol.Operator {
		case "+":
			result += operand
		case "*":
			result *= operand
		}
	}
	return result
}
