package less_than

import (
	"../instruction"
)

const ParametersCount int = 3

type Instruction int

func (i Instruction) Execute(memory []int, ip int) {
	if ip+i.GetParametersCount()+1 >= len(memory) {
		return
	}

	firstValue := instruction.GetSourceParameterValue(memory, ip, int(i), 0)
	secondValue := instruction.GetSourceParameterValue(memory, ip, int(i), 1)
	resultIndex := instruction.GetDestinationParameterValue(memory, ip, 2)

	if firstValue < secondValue {
		memory[resultIndex] = 1
	} else {
		memory[resultIndex] = 0
	}
}

func (i Instruction) GetParametersCount() int {
	return ParametersCount
}
