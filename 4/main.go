package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFileName string = "input.txt"

type Passport struct {
	Byr string //(Birth Year)
	Iyr string //(Issue Year)
	Eyr string //(Expiration Year)
	Hgt string //(Height)
	Hcl string //(Hair Color)
	Ecl string //(Eye Color)
	Pid string //(Passport ID)
	Cid string //(Country ID)
}

func main() {
	passportList := readInput(inputFileName)
	fmt.Println(passportList)

	validPassCount := 0
	for _, passport := range passportList {
		fmt.Println(*passport)
		if passport.ValidateFields() {
			validPassCount++
		}
	}

	fmt.Println("Valid passports: ", validPassCount)
}

func (p *Passport) ValidateFields() bool {
	if p.Byr == "" {
		return false
	}
	if p.Iyr == "" {
		return false
	}
	if p.Eyr == "" {
		return false
	}
	if p.Hgt == "" {
		return false
	}
	if p.Hcl == "" {
		return false
	}
	if p.Ecl == "" {
		return false
	}
	if p.Pid == "" {
		return false
	}


	return true
}

// ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
// byr:1937 iyr:2017 cid:147 hgt:183cm
func (p *Passport) ParseLine(line string) {
	blocks := strings.Split(line, " ")

	for _, block := range blocks {
		pair := strings.Split(block, ":")
		switch pair[0] {
		case "byr":
			p.Byr = pair[1]
		case "iyr":
			p.Iyr = pair[1]
		case "eyr":
			p.Eyr = pair[1]
		case "hgt":
			p.Hgt = pair[1]
		case "hcl":
			p.Hcl = pair[1]
		case "ecl":
			p.Ecl = pair[1]
		case "pid":
			p.Pid = pair[1]
		case "cid":
			p.Cid = pair[1]
		}
	}
}


func readInput(fileName string) []*Passport {
	inputFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Unable to open file")
	}

	scanner := bufio.NewScanner(inputFile)

	scanner.Split(bufio.ScanLines)
	passportList := make([]*Passport, 0)
	head := &Passport{}
	passportList = append(passportList, head)

	for scanner.Scan() {
		line := scanner.Text()
		head.ParseLine(line)

		if line == "" {
			head = &Passport{}
			passportList = append(passportList, head)
		}
	}

	err = inputFile.Close()
	if err != nil {
		log.Fatal("Error closing file")
	}

	return passportList
}

func must(err error) {
	if err != nil {
		panic("Something went wrong")
	}
}