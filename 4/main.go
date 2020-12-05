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

type Passport struct {
	Byr string //(Birth Year)
	Iyr string //(Issue Year)
	Eyr string //(Expiration Year)
	Hgt string //(Height)
	Hcl string //(Hair Color)
	Ecl string //(Eye Color)
	Pid string //(Passport ID)
	Cid string //(Country ID)
	IsValid bool
}

func main() {
	passportList := readInput(inputFileName)
	fmt.Println(passportList)

	validPassCount := 0
	for _, passport := range passportList {
		if passport.ValidateAll() {
			passport.IsValid = true
			validPassCount++
		} else {
			passport.IsValid = false
		}

		if passport.IsValid == true {
			fmt.Println("\033[32m", *passport)
		} else {
			fmt.Println("\033[31m", *passport)
		}

	}

	fmt.Println("\033[0m", "Valid passports: ", validPassCount)
}

// four digits; at least 1920 and at most 2002
func (p *Passport) ValidateByr() bool {
	if len(p.Byr) != 4 {
		return false
	}

	digits, err := strconv.Atoi(p.Byr)
	if err != nil {
		return false
	}

	if digits < 1920 || digits > 2002 {
		return false
	}

	return true
}

// four digits; at least 2010 and at most 2020
func (p *Passport) ValidateIyr() bool {
	if len(p.Iyr) != 4 {
		return false
	}

	digits, err := strconv.Atoi(p.Iyr)
	if err != nil {
		return false
	}

	if digits < 2010 || digits > 2020 {
		return false
	}

	return true
}

// four digits; at least 2020 and at most 2030.
func (p *Passport) ValidateEyr() bool {
	if len(p.Eyr) != 4 {
		return false
	}

	digits, err := strconv.Atoi(p.Eyr)
	if err != nil {
		return false
	}

	if digits < 2020 || digits > 2030 {
		return false
	}

	return true
}

// hgt (Height) - a number followed by either cm or in:
//   If cm, the number must be at least 150 and at most 193.
//   If in, the number must be at least 59 and at most 76.
func (p *Passport) ValidateHgt() bool {
	if strings.HasSuffix(p.Hgt, "cm") {
		number, err := strconv.Atoi(strings.TrimSuffix(p.Hgt, "cm"))
		if err != nil {
			return false
		}
		if number < 150 || number > 193 {
			return false
		}
		return true
	}
	if strings.HasSuffix(p.Hgt, "in") {
		number, err := strconv.Atoi(strings.TrimSuffix(p.Hgt, "in"))
		if err != nil {
			return false
		}
		if number < 59 || number > 76 {
			return false
		}
		return true
	}

	return false
}

// a # followed by exactly six characters 0-9 or a-f.
func (p *Passport) ValidateHcl() bool {
	if len(p.Hcl) != 7 {
		return false
	}
	if strings.HasPrefix(p.Hcl, "#") {
		colour := strings.TrimPrefix(p.Hcl, "#")
		if isAlphaNum(colour) {
			return true
		}
	}

	return false
}

func isAlphaNum(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'f') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

// exactly one of: amb blu brn gry grn hzl oth.
func (p *Passport) ValidateEcl() bool {
	switch p.Ecl {
	case "amb":
		return true
	case "blu":
		return true
	case "brn":
		return true
	case "gry":
		return true
	case "grn":
		return true
	case "hzl":
		return true
	case "oth":
		return true
	}

	return false
}

// a nine-digit number, including leading zeroes.
func (p *Passport) ValidatePid() bool {
	if len(p.Pid) == 9 {
		_, err := strconv.Atoi(p.Pid)
		if err != nil {
			return false
		}

		return true
	}

	return false
}

// ValidateAll validates all
func (p *Passport) ValidateAll() bool {
	if !p.ValidateFieldsExist() {
		return false
	}
	if p.ValidateByr() && p.ValidateEcl() && p.ValidateEyr() && p.ValidateHcl() && p.ValidateHgt() && p.ValidateIyr() && p.ValidateIyr() && p.ValidatePid() {
		return true
	}

	return false
}

func (p *Passport) ValidateFieldsExist() bool {
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
