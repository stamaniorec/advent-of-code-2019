package instruction

type Instruction interface {
	Execute(memory []int, ip int, rb int)
	GetParametersCount() int
}

type ParameterMode int

const (
	PositionalMode ParameterMode = iota
	ImmediateMode
	RelativeMode
)

func GetParameterMode(instructionCode, operandIndex int) ParameterMode {
	instructionCode /= 100

	var i int
	for instructionCode > 0 && i < operandIndex {
		instructionCode /= 10
		i++
	}

	instructionCode = instructionCode % 10
	if instructionCode == 0 {
		return PositionalMode
	} else if instructionCode == 1 {
		return ImmediateMode
	} else {
		return RelativeMode
	}
}

func GetSourceParameterValue(memory []int, ip int, relativeBase int, instructionCode int, parameterIndex int) int {
	index := ip + parameterIndex + 1
	parameterMode := GetParameterMode(instructionCode, parameterIndex)
	if parameterMode == ImmediateMode {
		return memory[index]
	} else if parameterMode == RelativeMode {
		return memory[memory[index]+relativeBase]
	} else {
		return memory[memory[index]]
	}
}

func GetDestinationParameterValue(memory []int, ip int, relativeBase int, instructionCode int, parameterIndex int) int {
	index := ip + parameterIndex + 1
	parameterMode := GetParameterMode(instructionCode, parameterIndex)
	if parameterMode == PositionalMode {
		return memory[index]
	} else {
		return memory[index] + relativeBase
	}
}
