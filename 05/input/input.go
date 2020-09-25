package input

import (
	"fmt"

	"../instruction"
)

const ParametersCount int = 1

type Instruction int

func (i Instruction) Execute(memory []int, ip int) {
	if ip+i.GetParametersCount()+1 >= len(memory) {
		return
	}

	var consoleInput int
	// fmt.Println("Enter a value, please!")
	fmt.Scanln(&consoleInput)

	positionIndex := instruction.GetDestinationParameterValue(memory, ip, 0)
	memory[positionIndex] = consoleInput
}

func (i Instruction) GetParametersCount() int {
	return ParametersCount
}
