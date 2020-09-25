package adjust_relative_base

import (
	"../instruction"
)

const ParametersCount int = 1

type Instruction int

func (i Instruction) Execute(memory []int, ip int, rb int) int {
	if ip+i.GetParametersCount()+1 >= len(memory) {
		return rb
	}

	value := instruction.GetSourceParameterValue(memory, ip, rb, int(i), 0)
	return rb + value
}

func (i Instruction) GetParametersCount() int {
	return ParametersCount
}
