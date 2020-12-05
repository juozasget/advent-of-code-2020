package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
)

type BoardingPass struct {
	RowInst string
	ColInst string
	SeatID int
	colID int
	rowID int

}

const inputFileName string = "input.txt"
const rows int = 128
const cols int = 8

func main() {
	boardingPasses := readInput(inputFileName)
	highestSeatID := 0

	for _, boardingPass := range boardingPasses {
		boardingPass.ParseAll()
		boardingPass.SeatID = boardingPass.rowID * 8 + boardingPass.colID
		if boardingPass.SeatID > highestSeatID {
			highestSeatID = boardingPass.SeatID
		}

		fmt.Println(*boardingPass)
	}

	fmt.Println(highestSeatID)
}

func (bp *BoardingPass) ParseAll() {
	bp.ParseRows()
	bp.ParseCols()
}

func (bp *BoardingPass) ParseCols() {
	rowRange := cols
	lo := 0
	hi := 7
	for _, instruction := range bp.ColInst {
		rowRange = rowRange / 2
		if string(instruction) == "L" {
			hi -= rowRange
		} else if string(instruction) == "R" {
			lo += rowRange
		}
	}

	bp.colID = lo
}

func (bp *BoardingPass) ParseRows() {
	rowRange := rows
	lo := 0
	hi := 127
	for _, instruction := range bp.RowInst {
		rowRange = rowRange / 2
		if string(instruction) == "F" {
			hi -= rowRange
		} else if string(instruction) == "B" {
			lo += rowRange
		}
	}

	bp.rowID = lo
}

func readInput(fileName string) []*BoardingPass {
	inputFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Unable to open file")
	}

	scanner := bufio.NewScanner(inputFile)

	scanner.Split(bufio.ScanLines)
	boardingPasses := make([]*BoardingPass, 0)

	for scanner.Scan() {
		line := scanner.Text()
		newBoardingPass := &BoardingPass{
			RowInst: line[:7],
			ColInst: line[7:],
		}
		boardingPasses = append(boardingPasses, newBoardingPass)
	}

	err = inputFile.Close()
	if err != nil {
		log.Fatal("Error closing file")
	}

	return boardingPasses
}

func must(err error) {
	if err != nil {
		panic("Something went wrong")
	}
}