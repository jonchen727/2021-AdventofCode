package main

import (
	"encoding/csv"
	"fmt"

	//  "log"
	"os"
	"strconv"
	"strings"
	//  "reflect"
)

func fileinput() string {

	file := os.Args[1]

	if strings.Contains(file, ".csv") {
	} else {
		fmt.Println("Please select a csv file")
		os.Exit(1)
	}
	return file
}

func readcsv(file string) [][]string {
	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV")
	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	r.Comma = ' '

	csvLines, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	return csvLines
}

func calculate (records [][]string) (int, int, int) {
	var horizontal int
	var depth int
	var depth2 int
	var aim int
	for i := 0; i <len(records); i++ {
		command := records[i][0]
		value,_ := strconv.Atoi(records[i][1])
		
		if command == "forward" {
			horizontal += value
			depth2 += (aim * value)
		} else if command == "down" {
			depth += value
			aim += value
		} else if command == "up" {
			depth -= value
			aim -= value
		}
	}
	return horizontal,depth, depth2
}
func main() {
file := fileinput()
records := readcsv(file)
Horizontal, Depth, Depth2 := calculate(records)
fmt.Println("Horizontal:",Horizontal)
fmt.Println("Depth:", Depth)
fmt.Println("Answer:", (Horizontal*Depth))
fmt.Println("Answer 2:",(Depth2*Horizontal))


}