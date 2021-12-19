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

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	return csvLines
}

func increased (records [][]string) int{

  var prevline int
  counter:=0
  for i, line := range records {
    if i == 0 {
      newline, _  := strconv.Atoi(line[0])
      //fmt.Println(newline,"First Line")
    }else if i >= 1 {
      //fmt.Println("Prev", prevline)
      newline, _ := strconv.Atoi(line[0])
      //fmt.Println("New",newline)

      if (newline > prevline){
        //fmt.Println(line[0],"(increased)")
        counter+=1
      } else if (newline < prevline) {
        //fmt.Println(line[0],"(decreased)")
      } else if (prevline == newline) {
        //fmt.Println(line[0],"(same)")
      } else {
        //fmt.Println(line[0])
      }
    prevline=newline
    }
  }
  return counter
}


func slidingwindow (records [][]string) int{ 
  var previouswindow int
  var counter int

	for i := 0; i <= len(records)-3; i++ {
    A1, _ := strconv.Atoi(records[i][0])
		A2, _ := strconv.Atoi(records[i+1][0])
		A3, _ := strconv.Atoi(records[i+2][0])
    currentwindow := A1+A2+A3

    if i == 0 {
      //fmt.Println(currentwindow, "No Previous Sum")
      previouswindow = currentwindow
    } else {
        if (currentwindow > previouswindow) {
          //fmt.Println(currentwindow, "(increased)")
          counter += 1
        } else if (currentwindow < previouswindow) {
          //fmt.Println(currentwindow, "(decreased)")
        } else if (currentwindow == previouswindow) {
          //fmt.Println(currentwindow,"(same)")
        }
      previouswindow = currentwindow
      }
	}
  return counter
}


func main() {
	file := fileinput()
	records := readcsv(file)

  answer1 := increased(records)
  answer2 := slidingwindow(records)

  fmt.Println("Answer 1:",answer1)
  fmt.Println("Answer 2:",answer2)
}
