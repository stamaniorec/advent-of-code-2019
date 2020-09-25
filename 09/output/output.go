package output

import (
	"fmt"

	"../instruction"
)

const ParametersCount int = 1

type Instruction int

func (i Instruction) Execute(memory []int, ip int, rb int) {
	if ip+i.GetParametersCount()+1 >= len(memory) {
		return
	}

	value := instruction.GetSourceParameterValue(memory, ip, rb, int(i), 0)
	fmt.Println(value)
}

func (i Instruction) GetParametersCount() int {
	return ParametersCount
}
