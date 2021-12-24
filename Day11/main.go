package main

import (
	"fmt"
	//"image/color"
	"io/ioutil"
	"os"

	//	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"
	//"go.uber.org/zap/internal/color"
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
	//ilines := StringtoInt(lines)
	ilines := NestArray(lines)


	SimulateTurn(ilines, 100)
	




	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

//func StringtoInt(sLines []string) []int {
//	var iLines []int
//
//	for _, sLine := range sLines { //for each line in sLine
//		iLine, _ := strconv.Atoi(sLine) //convert the string to int
//		iLines = append(iLines, iLine)  //add converted line to iLines
//	}
//	return iLines
//}

func SimulateTurn (ilines [][]int, turns int) {
	var flashes int
	for i := 0 ; i <= turns; i++ {
//		turnflash := 0
		if i == 0 {
			fmt.Println("Turn = 0 ")
			for _,line := range ilines{
				//fmt.Println(line)
				for _,character := range line {
					Colorize(character)
				}
				fmt.Println()
			}
		}else { 
			for _,line := range ilines{
				for i := 0 ; i < len(line); i++ {
					line[i] += 1
					}
				}
				for {
					trigger := false 
					for j, lines := range ilines{
						for i := 0 ; i < len(lines); i++ {
							if lines[i] > 9 {
								//fmt.Println("we exploded")
								trigger = true
								lines[i] = -1000
								for q := -1 ; q <= 1 ; q++ {
									jq := q+j
									if jq >= 0 && jq < len(lines){
										for r := -1; r <= 1; r++ {
											ri := r+i
											if ri >= 0 && ri < len(ilines) {
												ilines[jq][ri] += 1
											}
										}
									}
									
								}
							}
						}
					}
					if trigger == false {
						for _,lines := range ilines{
							for i:=0 ; i< len(lines); i++ {
								if lines[i] < 0 {
									lines[i] = 0
									flashes += 1
								}
							}
						}
						break
					}
				}
				fmt.Println()
				fmt.Println("Turn =",i)
				for _,line := range ilines{
					//fmt.Println(line)
					for _,character := range line {
						Colorize(character)
					}
					fmt.Println()
			}
			}


		}
		fmt.Println()
		fmt.Println(flashes)

}
		



func NestArray (lines []string) [][]int {
	var Arrays [][]int 
	for _,lines := range lines {
		var Array []int
		
		for _,character := range lines {
			a,_ := strconv.Atoi(string(character))
			Array = append(Array, a)
		}
		Arrays = append(Arrays,	Array)
	}
	return Arrays
}

func Colorize (character int) {
	switch character {
	case 0: color.Bold.Print(character)
	default: fmt.Print(character)
	}
	
}