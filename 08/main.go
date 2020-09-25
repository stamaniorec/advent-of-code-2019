package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

const inputFileName string = "input.txt"

func readImageDataFromFile() (string, error) {
	dat, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		return "", fmt.Errorf("Can not open input file: %w", err)
	}

	return strings.TrimSpace(string(dat)), nil
}

func GetLayerWithFewestZeroes(imageData string, width, height int) int {
	imageSize := width * height

	var (
		zeroesCount            int
		layerIndex             int
		fewestZeroesLayerIndex int
	)
	minZeroesCount := math.MaxInt32

	for i, pixel := range imageData {
		if pixel == rune('0') {
			zeroesCount++
		}

		isLastOfLayer := (i+1)%imageSize == 0
		if isLastOfLayer {
			if zeroesCount < minZeroesCount {
				minZeroesCount = zeroesCount
				fewestZeroesLayerIndex = layerIndex
			}

			zeroesCount = 0
			layerIndex++
		}
	}

	return fewestZeroesLayerIndex
}

func isTransparent(pixel byte) bool {
	return pixel == "2"[0]
}

func printPixel(pixel byte) {
	switch pixel {
	case "0"[0]:
		fmt.Print(".")
	default:
		fmt.Print("#")
	}
}

func PrintDecodedImage(imageData string, width, height int) {
	imageSize := width * height

	for i := 0; i < imageSize; i++ {
		pixelIndex := i
		for pixelIndex < len(imageData) && isTransparent(imageData[pixelIndex]) {
			pixelIndex += imageSize
		}

		printPixel(imageData[pixelIndex])

		isLastOfLayer := (i+1)%width == 0
		if isLastOfLayer {
			fmt.Println()
		}
	}
}

func main() {
	imageData, err := readImageDataFromFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	width := 25
	height := 6

	imageSize := width * height

	layerIndex := GetLayerWithFewestZeroes(imageData, width, height)

	var onesCount int
	var twosCount int
	for i := layerIndex * imageSize; i < (layerIndex+1)*imageSize; i++ {
		switch imageData[i] {
		case "1"[0]:
			onesCount++
		case "2"[0]:
			twosCount++
		}
	}

	fmt.Println(onesCount * twosCount)
	fmt.Println()
	fmt.Println()

	PrintDecodedImage(imageData, width, height)
}
