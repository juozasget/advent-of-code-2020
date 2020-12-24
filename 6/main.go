package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFileName string = "input.txt"

func main() {
	groupList := readInput(inputFileName)
	count := 0
	// range over groups
	for i, group := range groupList {
		checkSet := make(map[string]int)
		var n int
		var member string
		// range over members of groups
		for n, member = range group {
			// range over chars of member answers
			for _, char := range member {
				checkSet[string(char)]++
			}
		}
		for key, value := range checkSet {
			fmt.Printf("Checking key: %v, value: %v if it's equal to %v\n", key, value, n+1)
			if value == n+1 {
				count++
			}
		}
		fmt.Printf("Group: %v Count: %v\n", i, count)
	}
	fmt.Println(count)
}

func readInput(fileName string) [][]string {
	inputFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Unable to open file")
	}

	scanner := bufio.NewScanner(inputFile)

	scanner.Split(bufio.ScanLines)
	groupList := make([][]string, 0)
	newGroup := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			newGroup = append(newGroup, line)
		} else {
			groupList = append(groupList, newGroup)
			newGroup = make([]string, 0)
		}
	}

	groupList = append(groupList, newGroup)

	err = inputFile.Close()
	if err != nil {
		log.Fatal("Error closing file")
	}

	return groupList
}