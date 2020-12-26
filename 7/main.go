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

func main() {
	bagRules := readInput(inputFileName)
	fmt.Println(bagRules)
	fmt.Println(len(bagRules))

	var counter int

	for bagName, _ := range bagRules {
		fmt.Printf("Checking %v\n", bagName)
		counterBefore := counter
		traversBagRules(bagRules, bagName, &counter)
		counterAfter := counter
		deltaCounter := counterAfter - counterBefore
		if deltaCounter > 1 {
			counter = counterBefore + 1
		}
	}

	fmt.Printf("Result: %v\n", counter)
}

func traversBagRules(bagRules map[string]map[string]int, bagName string, count *int) {
	_, exists := bagRules[bagName]["none"]

	if exists {
		fmt.Println("Reached end!")
		return
	}

	_, exists = bagRules[bagName]["shiny-gold"]
	if exists {
		*count++
		fmt.Printf("You can put shiny-gold bag in %v! Current count: %v\n", bagName, *count)
		return
	}

	for key, _ := range bagRules[bagName] {
		traversBagRules(bagRules, key, count)
	}
}

func readInput(fileName string) map[string]map[string]int{
	inputFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Unable to open file")
	}

	scanner := bufio.NewScanner(inputFile)
	bagList := make(map[string]map[string]int)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		unparsedBagRule := strings.Split(line, "contain")
		bagType := strings.Split(unparsedBagRule[0], " ")
		bagName := bagType[0] + "-" + bagType[1]

		parsedRule := make(map[string]int)

		if strings.TrimSpace(unparsedBagRule[1]) == "no other bags." {
			parsedRule["none"] = 0
		} else {
			rules := strings.Split(strings.TrimSpace(unparsedBagRule[1]), ",")
			for _, rule := range rules {
				bagType := strings.Split(strings.TrimSpace(rule), " ")
				bagRule := bagType[1] + "-" + bagType[2]
				bagRuleCount, _ := strconv.Atoi(bagType[0])
				parsedRule[bagRule] = bagRuleCount
			}
		}

		bagList[bagName] = parsedRule
	}

	err = inputFile.Close()
	if err != nil {
		log.Fatal("Error closing file")
	}

	return bagList
}