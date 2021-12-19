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

	csvLines, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	return csvLines
}

func counter (records [][]string, index int) (int,int){
	var count0 int
	var count1 int
	for j := 0 ; j < len(records) ; j++ {
		character := string(([]rune(records[j][0]))[index])
		if character == "0" {
			//fmt.Println("its a 0")
			count0 += 1
		} else if character == "1" {
			//fmt.Println("its a 1")	
			count1 += 1
    	}
	}
	return count0, count1
}

func stringmaker (records [][]string) (string, string) {
	var gamma string
	var epsilon string

	for i := 0 ; i < len(records[0][0]) ; i++ {
		count0, count1 := counter(records, i)
		if (count0 > count1) {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
		
	}
	return gamma, epsilon 
}


func main() {
file := fileinput()
records := readcsv(file)
gamma, epsilon := stringmaker(records)
gammanum,_ := strconv.ParseInt(gamma,2,64)
epsilonnum,_ := strconv.ParseInt(epsilon,2,64)
fmt.Println("Gamma Binary:",gamma)
fmt.Println("Gamma:",gammanum)
fmt.Println("Epsilon Binary:",epsilon)
fmt.Println("Epsilon:",epsilonnum)
fmt.Println("Power Consumption:",(gammanum*epsilonnum))
}