package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"math"
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
	ilines := StringtoInt(lines)


	sd, average := MathStuff(ilines)

	fmt.Printf ("##### Data Set Information #####\nAverage: %.0f \nStandard Deviation: %.0f \n", average, sd)


	startpoint := (int(average) - (int(sd)/3))
	
	Fuel1, Iterations1 :=  Answer1(sd, startpoint, ilines)
	fmt.Printf("\nAnswer 1: %d fuel \nThis took %d Iterations\n", Fuel1,Iterations1)

	
	Fuel2, Iterations2 :=  Answer2(sd, startpoint, ilines)
	fmt.Printf("\nAnswer 2: %d fuel \nThis took %d Iterations\n", Fuel2,Iterations2)




	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func Answer1(sd float64, startpoint int, ilines []int) (int, int) {
	var fuel float64 
	var fuelprev float64
	var iterations int
	for i := 0; i <= int(sd); i++ {
		fuel = 0
		if i == 0 {
			
			for _, line := range ilines {
				fuel += math.Abs(float64(startpoint - line))
			}
			fuelprev = fuel
		}else if i >= 1 {
			for _, line := range ilines {
				fuel += math.Abs(float64(startpoint - line))
		}
			if fuel > fuelprev { 
				iterations = i
				fuel=fuelprev 
				break
					
			} else {
				fuelprev = fuel
			}
		}
		startpoint += 1
		//fmt.Println(fuel)
	}
	return int(fuel),iterations
}		

func Answer2(sd float64, startpoint int, ilines []int) (int, int) {
	var fuel float64 
	var fuelprev float64
	var iterations int
	for i := 0; i <= int(sd); i++ {
		//fmt.Println("Iteration:",i)
		fuel = 0
		if i == 0 {
			for _,line := range ilines {
				var distance int
				distance = int(math.Abs(float64(startpoint - line)))
				for j:=1 ; j <= distance; j++ {
					fuel += float64(j)
				}	
			} 
			fuelprev = fuel
		}else if i >= 1 {
			for _, line := range ilines {
				var distance int
				distance = int(math.Abs(float64(startpoint - line)))
				for j:=1 ; j <= distance; j++ {
					fuel += float64(j)
		}
	}
			if fuel > fuelprev { 
				iterations = i
				fuel=fuelprev 
				break
					
			} else {
				fuelprev = fuel
			}
		}
		startpoint += 1
		//fmt.Println(fuel)
	}
	return int(fuel), iterations
}	



func MathStuff(ilines []int) ( float64, float64) {
	var sd float64
	var sum float64
	for _, line := range ilines {
		sum += float64(line)
	}

	average := (float64(sum)) / float64(len(ilines))
	
	for _, line := range ilines {
		sd += math.Pow(float64(line)-average, 2)
	}

	sd = math.Sqrt(sd / float64(len(ilines)))

	return sd, average
}



