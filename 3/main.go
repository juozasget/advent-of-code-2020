package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
)

const inputFileName string = "input.txt"

type Tile struct {
	X int
	Y int
	TileType string
}

func main() {
	landMap := readInput(inputFileName)
	for _, line := range landMap {
		fmt.Println(line)
	}

	mapWidth := len(landMap[0])
	mapLength := len(landMap)
	fmt.Println("Width: ", mapWidth)
	fmt.Println("Length: ", mapLength)

	for i:= 0; i < (mapLength * 3) / 31 + 31; i++ {
		extendMap(landMap)
	}
}

func extendMap(line string) string {
	return ""
}

func readInput(fileName string) []string {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal("Unable to open file")
	}

	scanner := bufio.NewScanner(inputFile)

	scanner.Split(bufio.ScanLines)
	var input []string

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	err = inputFile.Close()
	if err != nil {
		log.Fatal("Error closing file")
	}

	input = append(input, "")

	for i:= 0; i < len(input[0]); i++ {
		input[len(input)-1] += "X"
	}

	return input
}

func must(err error) {
	if err != nil {
		panic("Something went wrong")
	}
}