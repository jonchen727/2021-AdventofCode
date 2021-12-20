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

func counter(records [][]string, index int) string {
	var count0 int
	var count1 int
	var result string
	for j := 0; j < len(records); j++ {
		character := string(([]rune(records[j][0]))[index])
		if character == "0" {
			//fmt.Println("its a 0")
			count0 += 1
		} else if character == "1" {
			//fmt.Println("its a 1")
			count1 += 1
		}
	}
	if count0 > count1 {
		result = "0"
	} else if count0 < count1 {
		result = "1"
	} else if count0 == count1 {
		result = "same"
	}
	return result
}

func stringmaker(records [][]string) (string, string) {
	var gamma string
	var epsilon string

	for i := 0; i < len(records[0][0]); i++ {
		result := counter(records, i)
		if result == "0" {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}

	}
	return gamma, epsilon
}

func co2scrub(records [][]string) string {
	co2record := records
	for i := 0; i < len(records[0][0]); i++ {
		result := counter(co2record, i)
		//fmt.Println("Result:", result)
		for j := 0; i < len(records); j++ {

			ctr := len(co2record)
			fmt.Println(j, ctr)
			if (j < ctr-1)  {

				character := string(([]rune(co2record[j][0]))[i])
				fmt.Println("Index:", i, "Line:", j, "Character:", character)
				fmt.Println(co2record)
				if result == character {
					//fmt.Println("case 1")
					co2record = append(co2record[:j], co2record[j+1:]...)
					j -= 1
					ctr -= 1
					//co2record2 = append(oxyge)
				} else if result == "same" && character == "1" {
					//fmt.Println("case 2")
					co2record = append(co2record[:j], co2record[j+1:]...)
					j -= 1
					ctr -= 1
				} else if result == "same" && character == "0" {
					//fmt.Println("case 3")
				} else if result != character {
					//fmt.Println("case 4")
				}
			} else {
				//fmt.Println("done")
				break
			}
		}
	}
	return co2record[0][0]
}
func oxygengen(records [][]string) string {
	oxygenrecord := records
	fmt.Println(oxygenrecord)
	for i := 0; i < len(records[0][0]); i++ {
		result := counter(oxygenrecord, i)
		//fmt.Println("Result:",result)
		for j := 0; i < len(records); j++ {
			ctr := len(oxygenrecord)
			fmt.Println(j,ctr)
			if j < ctr {
				character := string(([]rune(oxygenrecord[j][0]))[i])
				//fmt.Println("Index:",i,"Line:",j,"Character:", character,"Result:",result)
				//mt.Println(oxygenrecord)
				if result == character {
					//fmt.Println("case 1")
					//oxygenrecord2 = append(oxyge)
				} else if result == "same" && character == "1" {
					//fmt.Println("case 2")
				} else if result == "same" && character == "0" {
					//fmt.Println("case 3")
					oxygenrecord = append(oxygenrecord[:j], oxygenrecord[j+1:]...)
					j -= 1
					ctr -= 1
				} else if result != character {
					//fmt.Println("case 4")
					oxygenrecord = append(oxygenrecord[:j], oxygenrecord[j+1:]...)
					j -= 1
					ctr -= 1
				}
			} else {
				//fmt.Println("done")
				break
			}
		}
	}
	return oxygenrecord[0][0]
}

func main() {
	file := fileinput()
	records := readcsv(file)
	fmt.Println(records)
	gamma, epsilon := stringmaker(records)
	gammanum, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonnum, _ := strconv.ParseInt(epsilon, 2, 64)
	oxygenrecords := records
	co2records := records

	fmt.Println(records)
	oxygengen := oxygengen(oxygenrecords)
	co2scrub := co2scrub(co2records)
	fmt.Println(records)
	
	fmt.Println(records)
	oxygengennum, _ := strconv.ParseInt(oxygengen, 2, 64)
	co2scrubnum, _ := strconv.ParseInt(co2scrub, 2, 64)
	fmt.Println("Gamma Binary:", gamma)
	fmt.Println("Gamma:", gammanum)
	fmt.Println("Epsilon Binary:", epsilon)
	fmt.Println("Epsilon:", epsilonnum)
	fmt.Println("Power Consumption:", (gammanum * epsilonnum))
	fmt.Println("Oxygen Gen Binary:", oxygengen)
	fmt.Println("Oxygen Gen:", oxygengennum)
	fmt.Println("CO2 Scrub Binary:", co2scrub)
	fmt.Println("CO2 Scrub:", co2scrubnum)
	fmt.Println("Answer2:",(co2scrubnum*oxygengennum))
	//co2record := records

}
