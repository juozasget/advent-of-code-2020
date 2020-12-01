package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFileName string = "input.txt"

func main() {
	input := getInput()

	for i := 0; i < len(input); i++ {
		for n := 0; n < len(input); n++ {
			if n == i {
				break
			}
			if input[i]+input[n] == 2020 {
				fmt.Println("Answer: ", input[i] * input[n])
			}
		}
	}
}

func getInput() []int {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal("Unable to open file")
	}

	scanner := bufio.NewScanner(inputFile)

	scanner.Split(bufio.ScanLines)
	var input []int

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Error converting from string to integer")
		}
		input = append(input, number)
	}

	err = inputFile.Close()
	if err != nil {
		log.Fatal("Errro closing file")
	}

	return input
}