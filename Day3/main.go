package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func FiletoArray(delim string) []string {
	var lines []string

	if len(os.Args) > 1 { // if file argument is provided
		file := os.Args[1]                  //takes 1st arg as file name
		if strings.Contains(file, ".txt") { //checks if file is .txt
			bytes1, _ := ioutil.ReadFile(file)    //read file convert to bytes
			input := string(bytes1)               //convert bytes to string
			lines = strings.Split((input), delim) //convert string to []string with function input as delimiter
		} else { // exit for non text file input
			fmt.Println("Please select a text file") //exits if not .txt file
			os.Exit(69)
		}
	} else { // exit for no argument input
		fmt.Println("Cannot run command with no file.")
		os.Exit(420)
	}

	return lines //returns final []string
}

func PowerConsumption(lines []string) (string, int, string, int, int) {
	sums := make([]int, len(lines[0])) //creates an int array with number of characters on a single line

	for _, line := range lines { //iterates over each line
		for j, character := range line { //iterates over each index
			if string(character) == "1" {
				sums[j] += 1 //adds 1 to sum array with index equal to j this gives number of 1's per index summed per row
			}
		}
	}
	var gammab string
	gamma := 0
	var epsilonb string
	epsilon := 0

	for j, sum := range sums { // will interate up from 0 to range of sums (5)
		index := len(sums) - j - 1 //calculates index by subtracting 5-1 to get index 0,4 and then subtracts iterative count binary is reverse ordered
		if sum > len(lines)/2 {    //if 1's are greater
			gamma += int(math.Exp2(float64(index))) //uses exp2 on index to convert binary to dec
			gammab += "1"                           //if 1s are greater add 1 to gamma binary string
			epsilonb += "0"                         //if 0's are lesser add 1 to epsilon binary string
		} else { //uses inverse case where 1's are less to generate a 5 bit string
			epsilon += int(math.Exp2(float64(index)))
			gammab += "0"   //if 0's are greater add 0 to gamma binary string
			epsilonb += "1" //if 1's are lesser add 1 to epsilon binary string
		}
	}
	return gammab, gamma, epsilonb, epsilon, (gamma * epsilon)
}

func GasSystem(lines []string, mode string) (int64, string) {
	var gas int64
	var gas_b string
	//var CO2 int64
	//var CO2_b string
	index := 0
	sum := 0
	currentLines := lines
	var Lines_0 []string
	var Lines_1 []string

	for { //infinite for loop
		for _, line := range currentLines { //iterate for lines in set of current lines
			if string(line[index]) == "1" { //since this is with in a for loop each instances is type string so index can be used
				sum += 1                        // counter for number of 1's in index
				Lines_1 = append(Lines_1, line) // adds current line to array
			} else {
				Lines_0 = append(Lines_0, line) // adds current line to array
			}
		}

		index += 1 //increase index for next loop
		if mode == "O2" {
			if len(Lines_1) >= len(Lines_0) { //if there are GTE 1's at index position than 0's
				currentLines = Lines_1 //set currentLines to lines with most index position 1's
			} else {
				currentLines = Lines_0 //else set currentLines to lines with most index position 0's
			}
		} else if mode == "CO2" {
			if len(Lines_1) >= len(Lines_0) { //if there are GTE 1's at index position than 0's
				currentLines = Lines_0 //set currentLines to lines with least index position 0's
			} else {
				currentLines = Lines_1 //else set currentLines to lines with least index position 1's
			}
		}

		Lines_0 = nil
		Lines_1 = nil

		if len(currentLines) == 1 { //when there is one line left
			gas_b = currentLines[0]                 //set oxygen binary to last line
			gas, _ = strconv.ParseInt(gas_b, 2, 64) //convert binary to base64 dec
			break
		}
	}
	return gas, gas_b
}

func LifeSupport(lines []string) (int64, string, int64, string, int64) {
	O2, O2_b := GasSystem(lines, "O2")
	CO2, CO2_b := GasSystem(lines, "CO2")

	return O2, O2_b, CO2, CO2_b, (O2 * CO2)
}

func main() {
	start := time.Now() //sets current time to start time

	lines := FiletoArray("\n")

	gammab, gamma, epsilonb, epsilon, powerconsumption := PowerConsumption(lines)

	fmt.Println("Gamma:", gamma, "Gamma Binary:", gammab)
	fmt.Println("Epsilon:", epsilon, "Epsilon Binary:", epsilonb)
	fmt.Println("Power Consumption:", powerconsumption)

	O2, O2_b, CO2, CO2_b, Rating := LifeSupport(lines)

	fmt.Println()
	fmt.Println("Oxygen Generator Rating:", O2, "Oxygen Generator Rating Binary:", O2_b)
	fmt.Println("CO2 Scrubber Rating:", CO2, "CO2 Scrubber Rating Binary:", CO2_b)
	fmt.Println("Life Support Rating:", Rating)

	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}
