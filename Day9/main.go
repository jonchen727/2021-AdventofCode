package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
func StringtoInt(sLines []string) []int {
	var iLines []int

	for _, sLine := range sLines { //for each line in sLine
		iLine, _ := strconv.Atoi(sLine) //convert the string to int
		iLines = append(iLines, iLine)  //add converted line to iLines
	}
	return iLines
}

func main() {
	start := time.Now() //sets current time to start time
	lines := (FiletoArray("\n", 1))
	//ilines := StringtoInt(lines)

	


	fmt.Println("Answer 1:",Answer1(lines))

	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func Answer1(lines []string) int {
	var sum int
	for l, line := range lines {

		for i, index := range line {
			var right string
			var left string
			var top string
			var bottom string

			var neighbors float64
			var lower float64

			x := string(index)
			ix, _ := strconv.Atoi(x)
			if i < len(string(line))-1 {
				right = string(lines[l][i+1])

				iright, _ := strconv.Atoi(right)

				neighbors += 1
				if iright > ix {
					lower += 1
				}

			}
			if i > 0 {
				left = string(lines[l][i-1])
				ileft, _ := strconv.Atoi(left)
				neighbors += 1
				if ileft > ix {
					lower += 1
				}

			}

			if l < len(lines)-1 {
				bottom = string(lines[l+1][i])
				ibottom, _ := strconv.Atoi(bottom)
				neighbors += 1
				if ibottom > ix {
					lower += 1
				}

			}

			if l > 0 {
				top = string(lines[l-1][i])
				itop, _ := strconv.Atoi(top)
				neighbors += 1
				if itop > ix {
					lower += 1
				}

			}

			if neighbors == lower {
				sum += (ix + 1)
			}
		}

	}

	return sum
}

