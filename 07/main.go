package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"

	"./amplifier"
	"./permutation"
)

const inputFileName string = "input.txt"

func readAmplifierProgramFromFile() (string, error) {
	dat, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		return "", fmt.Errorf("Can not open input file: %w", err)
	}

	return strings.TrimSpace(string(dat)), nil
}

func GetOutputSignal(phaseSettingConfig []int, amplifierProgram string) (int, error) {
	var inputSignal int

	for _, phaseSetting := range phaseSettingConfig {
		outputSignal, err := amplifier.Run(amplifierProgram, phaseSetting, inputSignal)
		if err != nil {
			return 0, err
		}

		inputSignal = outputSignal
	}

	return inputSignal, nil
}

func GetOutputSignalWithFeedback(phaseSettingConfig []int, amplifierProgram string) (int, error) {
	var inputs []chan int
	var outputs []chan int

	for _, phaseSetting := range phaseSettingConfig {
		amplifierInput := make(chan int, 1)
		amplifierOutput := make(chan int, 1)

		amplifierInput <- phaseSetting

		inputs = append(inputs, amplifierInput)
		outputs = append(outputs, amplifierOutput)

		go amplifier.RunOnGoing(amplifierProgram, amplifierInput, amplifierOutput)
	}

	var inputSignal int

	var i int
	for {
		inputs[i] <- inputSignal

		outputSignal, isRunning := <-outputs[i]
		if isRunning {
			inputSignal = outputSignal
		} else if i == len(phaseSettingConfig)-1 {
			break
		}

		i++
		if i >= len(phaseSettingConfig) {
			i = 0
		}
	}

	return inputSignal, nil
}

func main() {
	program, err := readAmplifierProgramFromFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	// phaseSettingOptions := []int{0, 1, 2, 3, 4}
	phaseSettingOptions := []int{5, 6, 7, 8, 9}
	maxOutputSignal := math.MinInt32

	phaseSettingConfigurations := permutation.GeneratePermutations(phaseSettingOptions)

	for _, config := range phaseSettingConfigurations {
		// outputSignal, err := GetOutputSignal(config, program)
		outputSignal, err := GetOutputSignalWithFeedback(config, program)
		if err != nil {
			fmt.Println(err)
			return
		}

		if outputSignal > maxOutputSignal {
			maxOutputSignal = outputSignal
		}
	}

	fmt.Println(maxOutputSignal)
}
