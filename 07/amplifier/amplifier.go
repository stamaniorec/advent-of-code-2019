package amplifier

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

func Run(program string, phaseSetting int, inputSignal int) (int, error) {
	cmd := exec.Command("../05/main", program)
	cmd.Stdin = strings.NewReader(fmt.Sprintf("%d\n%d\n", phaseSetting, inputSignal))

	cmdOutput, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	outputSignal, err := strconv.Atoi(strings.TrimSpace(string(cmdOutput)))
	if err != nil {
		return 0, err
	}

	return outputSignal, nil
}

func RunOnGoing(program string, inputChannel <-chan int, outputChannel chan<- int) error {
	cmd := exec.Command("../05/main", program)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	phaseSetting := <-inputChannel
	io.WriteString(stdin, fmt.Sprintf("%d\n", phaseSetting))

	inputSignal := <-inputChannel
	io.WriteString(stdin, fmt.Sprintf("%d\n", inputSignal))

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		lineOfOutput := scanner.Text()
		outputSignal, err := strconv.Atoi(lineOfOutput)
		if err != nil {
			return fmt.Errorf("Can not convert %s to number: %w", lineOfOutput, err)
		}

		outputChannel <- outputSignal

		inputSignal = <-inputChannel
		io.WriteString(stdin, fmt.Sprintf("%d\n", inputSignal))
	}

	close(outputChannel)

	return cmd.Wait()
}
