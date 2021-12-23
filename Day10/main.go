package main

import (
//	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"

	//	"sort"
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

	var Answer1  int
	for _,line := range lines {
		var w,x,y,z int // ) ] } > 
		var match string
		for _,character := range line {
			a := string(character)
			switch a {
				case "(":
					w += 1
					match += ")"
				case ")":
					w -= 1
				case "[":
					x += 1
					match += "]"
				case "]":
					x -= 1
				case "{":
					y += 1
					match += "}"
				case "}":
					y -= 1
				case "<": 
					z += 1
					match += ">"
				case ">":
					z -= 1
			} 
			if a == string(match[len(match)-1]) {
			match = match[:len(match)-1]
			} else if (a == "]" || a == "}" || a == ">" || a == ")") && (a != string(match[len(match)-1])) {
				fmt.Println(a)
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
			break
			}
		
		}
	}
	fmt.Println(Answer1)
	

	
	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

