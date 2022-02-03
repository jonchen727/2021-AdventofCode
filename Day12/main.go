package main

import (
	"fmt"
	"unicode"
	//"image/color"
	"io/ioutil"
	"os"

	//	"sort"
	//	"strconv"
	"strings"
	"time"
	//	"github.com/gookit/color"
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

type Routes struct {
	route string
	big bool 
}
type Bulbasur struct {
	location string 
	routes []Routes
}


func main() {
	start := time.Now() //sets current time to start time
	lines := (FiletoArray("\n", 1))
	//ilines := StringtoInt(lines)
	fmt.Println(lines)
	var Lines [][]string
	var AllThatShit []Bulbasur

	for _,line := range lines{
		char := strings.Split(line, "-")
		Lines = append(Lines, char)


	}
	fmt.Println(Lines)

	for _,lines := range Lines {
		var ThatShit Bulbasur
		var route Routes
		var location string
		var routes []Routes

		location = lines[0]
		big := IsUpper(lines[1])
		route = Routes{route: lines[1],big: big}
		routes = append(routes, route)
		ThatShit = Bulbasur{location: location, routes: routes}
		AllThatShit = append(AllThatShit, ThatShit)	
			
		location = lines[1]
		routes = nil

		big = IsUpper(lines[0])
		route = Routes{route: lines[0],big: big}
		routes = append(routes, route)
		ThatShit = Bulbasur{location: location, routes: routes}
		AllThatShit = append(AllThatShit, ThatShit)	
		}

		fmt.Println(AllThatShit)
		for i := 0; i < len(AllThatShit); i ++{
			a := AllThatShit[i].location
				for j := 0; j < len(AllThatShit); j++ {
					if j != i && a == AllThatShit[j].location {
						AllThatShit[i].routes = append(AllThatShit[i].routes, AllThatShit[j].routes...)
					}
				} 
				
		}
		fmt.Println(AllThatShit)
	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func IsUpper(s string) bool {
    for _, r := range s {
        if !unicode.IsUpper(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}