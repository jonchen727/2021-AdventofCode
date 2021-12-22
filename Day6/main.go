package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	//	"math"
	"os"
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
	lines := (FiletoArray(",", 1))

	school1 := BreadFish(lines, 80)
	fmt.Println("Answer 1:", school1, "Fish")

	school2 := BreadFish(lines, 256)
	fmt.Println("Answer 2:", school2, "Fish")

	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func BreadFish(lines []string, days int) int {
	ilines := StringtoInt(lines)
	fishes := make([]int, 9)

	for _, line := range ilines {
		index := line
		fishes[index] += 1
	}
	sum := 0

	for day := 0; day <= days; day++ {
		sum = 0
		for _, count := range fishes {
			sum += count
		}
		//fmt.Println("Day:", day, fishes, sum)

		new_fishes := make([]int, 9)

		for index, count := range fishes {
			if index == 0 {
				new_fishes[6] += count
				new_fishes[8] += count
			} else {
				new_fishes[index-1] += count
			}
			fishes = new_fishes
		}

	}
	return sum
}
