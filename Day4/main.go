package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	//"strconv"
	"os"
	"strings"
	"time"
)
func FiletoArray(delim string, arg int) []string {
	var lines []string

	if len(os.Args) > arg { // if file argument is provided
		file := os.Args[arg]                  //takes 1st arg as file name
		if strings.Contains(file, ".txt") { //checks if file is .txt
			bytes, _ := ioutil.ReadFile(file)    //read file convert to bytes
			input := string(bytes)               //convert bytes to string
			lines = strings.Split((input), delim) //convert string to []string with function input as delimiter
		} else { // exit for non text file input
			fmt.Println("Please select a text file") //exits if not .txt file
			os.Exit(69)
		}
	} else { // exit for no argument input
		fmt.Println("Add more files to args")
		os.Exit(420)
	}
	return lines //returns final []string
}

func StringtoInt (sLines []string) []int {
	var iLines []int

	for _, sLine := range sLines { //for each line in sLine
		iLine, _ := strconv.Atoi(sLine) //convert the string to int
		iLines = append (iLines, iLine) //add converted line to iLines
	}
	return iLines
}

type Solution struct { //this struct will contain all variations of solutions
	numbers []int
}

type Board struct { //this struct will contain the boards and its solutions
	solutions []Solution
}

func GenerateSolutions (sBoards []string ) []Board {
	var stBoards []Board //create a variable of type Board struct array
	var numColumn int // number of columns
	var numRow int //number of rows
	
	for _, sBoard := range sBoards { //everything after this is operating on a per board basis
		var stBoard Board
		

		sRows := strings.Split(string(sBoard), "\n") //splits individual boards into a array of individual rows 
		numRow = len(sRows) //row count

		//adds all possible row solutions 

		for _, sRow := range sRows { //everything after is operating on a per row basis 
			iRow := StringtoInt(strings.Fields(sRow)) // takes each row and converts it into an array of individual integer numbers
			numColumn = len(iRow) // column count
			solution := Solution{numbers: iRow} // creates a new variable solution and stores iRows as a Solution struct
			stBoard.solutions = append(stBoard.solutions, solution) // appends solution to board.solutions struct for each row
			
		}

		for index := 0; index <= numColumn-1 ; index++ { //iterates through columns 
			solution := Solution{} //sets empty variable of type solution struct
			for line := 0 ; line <= numRow-1; line++ { //iterates through rows 
				solution.numbers = append(solution.numbers, stBoard.solutions[line].numbers[index])
			}
			stBoard.solutions = append(stBoard.solutions, solution) //add new solution to Board struct
		}
		stBoards = append(stBoards, stBoard) //adds the Board containging the solutions to the []Board  
	}
	return stBoards
}

func drawNumbers(draws []int, boards []Board) ([]Board, int, int) {
	for _, draw := range draws { //iterates for each draw 
		//fmt.Println(i, draw)
		for j, board := range boards { //iterates for each board 
			for k, solution := range board.solutions { //iterates for each solution
				for l, number := range solution.numbers {// iterates for each number 
					if draw == number { //if the draw matches the number
						numbers := solution.numbers 
						numbers[l] = numbers[len(numbers)-1] // move matched number to last index
						boards[j].solutions[k].numbers = numbers[:len(numbers)-1] //truncate last index 
						if len(boards[j].solutions[k].numbers) == 0 { //if a board has no numbers left 
							return boards, j, draw //return the final board state, board index, and draw number
						}
					}
				}
			}
		}
	}
	return make([]Board, 0), 0, 0 //if no winner is found return empty stuff used to stop error
}

func main() {
	start := time.Now() //sets current time to start time
	numbers := StringtoInt((FiletoArray(",",1)))
	boards := GenerateSolutions(FiletoArray("\n\n",2))
	finalboards, winningboard, lastdrawnumber := drawNumbers(numbers,boards)
	fmt.Println(finalboards[winningboard], lastdrawnumber)
	
	var sum int
	for i := 0 ; i <= 4; i++ {
		for _, number := range finalboards[winningboard].solutions[i].numbers {
		sum += number
		fmt.Println(number)
		}
	}
	fmt.Println(sum*lastdrawnumber)

	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}