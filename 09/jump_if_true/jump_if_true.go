package jump_if_true

import (
	"../instruction"
)

const ParametersCount int = 2

type Instruction int

func (i Instruction) Execute(memory []int, ip int, rb int) int {
	if ip+i.GetParametersCount()+1 >= len(memory) {
		return ip
	}

	condition := instruction.GetSourceParameterValue(memory, ip, rb, int(i), 0)
	newIp := instruction.GetSourceParameterValue(memory, ip, rb, int(i), 1)

	if condition != 0 {
		return newIp
	}

	return ip + i.GetParametersCount() + 1
}

func (i Instruction) GetParametersCount() int {
	return ParametersCount
}
