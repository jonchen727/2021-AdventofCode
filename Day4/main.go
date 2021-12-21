package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes1, _ := ioutil.ReadFile("data1.txt")
	bytes2, _ := ioutil.ReadFile("data2.txt")
	puzzleInput1 := string(bytes1)
	puzzleInput2 := string(bytes2)
	fmt.Println(createDraws(puzzleInput1))
	fmt.Println(createBoards(puzzleInput2)[1].rows)

}
func createDraws(puzzleInput1 string) []int {
	numStrings := strings.Split(string(puzzleInput1), ",")
	return convertToInts(numStrings)
}
func convertToInts(numStrings []string) []int {
	var numInts []int
	for _, numString := range numStrings {
		numInt, _ := strconv.Atoi(numString)
		numInts = append(numInts, numInt)
	}
	return numInts
}
type Row struct {
	numbers []int
}

type Board struct {
	rows []Row
}

func createBoards(puzzleInput2 string) []Board {
	var boards []Board
	boardsStrings := strings.Split(string(puzzleInput2), "\n\n")
	for _, boardStrings := range boardsStrings {
		var board Board

		// Add horizontal rows
		rowStrings := strings.Split(string(boardStrings), "\n")
		for _, rowString := range rowStrings {
			rowInts := convertToInts(strings.Fields(rowString))
			row := Row{numbers: rowInts}
			board.rows = append(board.rows, row)
		}

		// Add vertical rows
		for i := 0; i <= 4; i++ {
			row := Row{}
			for j := 0; j <= 4; j++ {
				row.numbers = append(row.numbers, board.rows[j].numbers[i])
			}
			board.rows = append(board.rows, row)
		}
		boards = append(boards, board)
	}
	return boards
}