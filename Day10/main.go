package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	//	"strconv"
	"strings"
	"time"
)

func FiletoArray(delim string, arg int) []string {
	var lines []string

	if len(os.Args) > arg { // if file argument is provided
		file := os.Args[arg]                //takes 1st arg as file name
		if strings.Contains(file, ".txt") { //checks if file is .txt
			bytes, _ := ioutil.ReadFile(file)     //read file convert to bytes
			input := string(bytes)                //convert bytes to string
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

func main() {
	start := time.Now() //sets current time to start time
	lines := (FiletoArray("\n", 1))

	Answer1, Answer2 := Answers(lines)
	fmt.Println("Answer 1:", Answer1)
	fmt.Println("Answer 2:", Answer2)

	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func Answers(lines []string) (int, int) {
	var Answer1 int
	var ScoreArray []int
	for _, line := range lines {
		var match string
		for _, character := range line {
			a := string(character)
			switch a {
			case "(":
				match += ")"
			case "[":
				match += "]"
			case "{":
				match += "}"
			case "<":
				match += ">"

			}
			if a == string(match[len(match)-1]) {
				match = match[:len(match)-1]
			} else if (a == "]" || a == "}" || a == ">" || a == ")") && (a != string(match[len(match)-1])) {

				switch a {
				case ")":
					Answer1 += 3
				case "]":
					Answer1 += 57
				case "}":
					Answer1 += 1197
				case ">":
					Answer1 += 25137

				}
				match = ""
				break
			}

		}
		var tpoints int

		for i := len(match) - 1; i >= 0; i-- {
			a := string(match[i])
			switch a {
			case ")":
				tpoints *= 5
				tpoints += 1
			case "]":
				tpoints *= 5
				tpoints += 2
			case "}":
				tpoints *= 5
				tpoints += 3
			case ">":
				tpoints *= 5
				tpoints += 4

			}

		}

		if tpoints > 0 {
			ScoreArray = append(ScoreArray, tpoints)
		}
	}

	sort.Ints(ScoreArray)
	//fmt.Println(ScoreArray)
	Answer2 := ScoreArray[(len(ScoreArray)-1)/2]
	return Answer1, Answer2
}
