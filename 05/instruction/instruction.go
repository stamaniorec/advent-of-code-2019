package instruction

type Instruction interface {
	Execute(memory []int, ip int)
	GetParametersCount() int
}

type ParameterMode int

const (
	PositionalMode ParameterMode = iota
	ImmediateMode
)

func GetParameterMode(instructionCode, operandIndex int) ParameterMode {
	instructionCode /= 100

	var i int
	for instructionCode > 0 && i < operandIndex {
		instructionCode /= 10
		i++
	}

	if instructionCode%10 == 0 {
		return PositionalMode
	} else {
		return ImmediateMode
	}
}

func GetSourceParameterValue(memory []int, ip int, instructionCode int, parameterIndex int) int {
	value := GetDestinationParameterValue(memory, ip, parameterIndex)

	if GetParameterMode(instructionCode, parameterIndex) == ImmediateMode {
		return value
	} else {
		return memory[value]
	}
}

func GetDestinationParameterValue(memory []int, ip int, parameterIndex int) int {
	return memory[ip+parameterIndex+1]
}
