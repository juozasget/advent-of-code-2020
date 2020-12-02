package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFileName string = "input.txt"

// 14-15 v: vdvvvvvsvvvvvfpv
type Password struct {
	Min int
	Max int
	Char string
	UserPass string
}

func main() {
	passwordSlice := readInput(inputFileName)

	validCount := 0
	for _, password := range passwordSlice {
		if string(password.UserPass[password.Min-1]) == password.Char && string(password.UserPass[password.Max-1]) != password.Char {
			validCount++
		}
		if string(password.UserPass[password.Max-1]) == password.Char && string(password.UserPass[password.Min-1]) != password.Char {
			validCount++
		}
	}

	fmt.Println("Valid password count: ", validCount)
}

func readInput(fileName string) []*Password {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal("Unable to open file")
	}

	scanner := bufio.NewScanner(inputFile)

	scanner.Split(bufio.ScanLines)
	var input []*Password

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, parseToPassword(line))
	}

	err = inputFile.Close()
	if err != nil {
		log.Fatal("Error closing file")
	}

	return input
}

// parseToPassword parses input line, example:
// 14-15 v: vdvvvvvsvvvvvfpv
func parseToPassword(line string) *Password {
	segments := strings.Split(line, " ")
	minMax := strings.Split(segments[0], "-")
	min, err := strconv.Atoi(minMax[0])
	must(err)
	max, err := strconv.Atoi(minMax[1])
	must(err)
	psw := Password{
		Min: min,
		Max: max,
		Char: strings.Trim(segments[1], ":"),
		UserPass: segments[2],
	}

	return &psw
}

func must(err error) {
	if err != nil {
		panic("Something went wrong")
	}
}