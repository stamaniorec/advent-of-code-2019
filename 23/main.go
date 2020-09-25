package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"sync"

	"./computer"
)

func readProgram() (string, error) {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return "", fmt.Errorf("Can not open input file: %w", err)
	}

	return strings.TrimSpace(string(dat)), nil
}

func isIdle(ins, outs []chan int) bool {
	for _, in := range ins {
		if len(in) > 0 {
			return false
		}
	}
	for _, out := range outs {
		if len(out) > 0 {
			return false
		}
	}
	return true
}

func main() {
	program, err := readProgram()
	if err != nil {
		fmt.Println(err)
		return
	}

	var inputChannels []chan int
	var outputChannels []chan int

	for addr := 0; addr < 50; addr++ {
		in := make(chan int, 1000)
		inputChannels = append(inputChannels, in)

		out := make(chan int, 1000)
		outputChannels = append(outputChannels, out)

		fmt.Printf("starting %d\n", addr)
		go computer.Run(program, addr, in, out)
	}

	var wg sync.WaitGroup

	var natX, natY int

	// NAT
	go func(x, y *int) {
		var dataPoints []bool
		const sustainSecs = 2

		for {
			point := isIdle(inputChannels, outputChannels)
			dataPoints = append(dataPoints, point)

			if len(dataPoints) >= sustainSecs {
				dataPoints = dataPoints[len(dataPoints)-sustainSecs:]

				isContinuouslyIdle := true
				for _, p := range dataPoints {
					if !p {
						isContinuouslyIdle = false
					}
				}

				if isContinuouslyIdle {
					inputChannels[0] <- *x
					inputChannels[0] <- *y

					fmt.Printf("sending NAT %d,%d\n", *x, *y)

					dataPoints = []bool{}
				}
			}

			time.Sleep(1*time.Second)
		}

	}(&natX, &natY)

	for addr := 0; addr < 50; addr++ {
		wg.Add(1)

		go func(srcAddr int, outs, ins []chan int) {
			defer wg.Done()

			out := outputChannels[srcAddr]

			for {
				if len(out) < 3 {
					time.Sleep(100)
					continue
				}

				destAddr := <-out
				x := <-out
				y := <-out

				if destAddr == 255 {
					natX = x
					natY = y
					continue

					// Part 1
					// fmt.Println("Done!")
					// fmt.Println(y)
					// return
				}

				// fmt.Printf("sending %d,%d from %d to %d\n", x, y, srcAddr, destAddr)

				inputChannels[destAddr] <- x
				inputChannels[destAddr] <- y
			}
		}(addr, outputChannels, inputChannels)
	}

	wg.Wait()
}
