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

	//for i:= 0; i < (mapLength * 3) / 31 + 31; i++ {
	//	extendMap(landMap)
	//}

	fmt.Println("Extend by: ", (mapLength * 3) / mapWidth + mapWidth)

	extendedMap := extendMap(landMap, (mapLength * 3) / mapWidth + mapWidth)


	for _, line := range extendedMap {
		fmt.Println(line)
	}

	cursorX := 0
	cursorY := 0

	treeCount := 0

	for {
		cursorX += 3
		cursorY += 1
		fmt.Println("X: ", cursorX)
		fmt.Println("Y: ", cursorY)
		if string(extendedMap[cursorY][cursorX]) == "#" {
			treeCount++
		}
		if string(extendedMap[cursorY][cursorX]) == "X" {
			break
		}
	}

	fmt.Println("Obstacle count: ", treeCount)
}

func extendMap(landMap []string, extendBy int) []string {
	newMap := make([]string, 0)

	for n, line := range landMap {
		newline := line
		for i:= 0; i < extendBy; i++ {
			newline += line
		}
		newMap = append(newMap, newline)
		fmt.Println("Line appended: ", n)
	}

	return newMap
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