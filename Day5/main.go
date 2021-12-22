package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"math"


	"os"
	"strings"
	//	"time"
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

type Point struct {
	x int
	y int
	hit int

}
type Coordinate struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

//type Move struct {
//	coordinates []Coordinate
//}


func main() {
	lines := (FiletoArray("\n",1))

	
	var stCoords []Coordinate

	stCoords = PopulateStruct(lines, stCoords)
	
	
	X_max,Y_max := CoordinateRange(stCoords)

	
	plane := CreatePlane(Y_max, X_max)

	PopulatePlane(stCoords, plane)

	fmt.Println("Answer 1:",Answer1(plane))

	//for _,plane := range plane {
	//	fmt.Println(plane)
	//}
}

func Answer1(plane [][]int) (int){
	var cnt int
	for _, y := range plane {
		for _, x := range y {
			if x > 1 {
				cnt += 1
			}
		}
	}
	return cnt
}

func PopulateStruct(lines []string, stCoords []Coordinate) []Coordinate {
	for _, line := range lines {
		sSet := strings.Split(string(line), " -> ")

		x1, _ := strconv.Atoi(strings.Split(string(sSet[0]), ",")[0])
		y1, _ := strconv.Atoi(strings.Split(string(sSet[0]), ",")[1])
		x2, _ := strconv.Atoi(strings.Split(string(sSet[1]), ",")[0])
		y2, _ := strconv.Atoi(strings.Split(string(sSet[1]), ",")[1])
		coordinate := Coordinate{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		}
		stCoords = append(stCoords, coordinate)

	}
	return stCoords
}

func CreatePlane(Y_max int, X_max int) [][]int {
	plane := make([][]int, Y_max+1)
	for i := range plane {
		plane[i] = make([]int, X_max+1)
	}
	fmt.Println("Len y", len(plane))
	fmt.Println("Len x", len(plane[0]))
	//fmt.Println(plane)
	return plane
}

func PopulatePlane(stCoords []Coordinate, plane [][]int) {
	for _, stCoord := range stCoords {

		if stCoord.x1 == stCoord.x2 {
			for i := 0; i <= int(math.Abs(float64(stCoord.y1)-float64(stCoord.y2))); i++ {
				x := stCoord.x1
				ymin := int(math.Min(float64(stCoord.y1), float64(stCoord.y2)))
				y := ymin + i
				//fmt.Println(x, y, 1)
				plane[y][x] += 1
			}
		} else if stCoord.y1 == stCoord.y2 {
			for i := 0; i <= int(math.Abs(float64(stCoord.x1)-float64(stCoord.x2))); i++ {
				y := stCoord.y1
				xmin := int(math.Min(float64(stCoord.x1), float64(stCoord.x2)))
				x := xmin + i
				//fmt.Println(x, y, 1)
				plane[y][x] += 1
			}
		}

	}
}

func CoordinateRange(stCoords []Coordinate) (int, int) {
	X_max := stCoords[0].x1
	Y_max := stCoords[0].y1
	
	for _, line := range stCoords {
		if line.x1 > X_max {
			X_max = line.x1
		} else if line.x2 > X_max {
			X_max = line.x2
		}
		if line.y1 > Y_max {
			Y_max = line.y1
		} else if line.y2 > Y_max {
			Y_max = line.y2
		}
	}
	
	return X_max, Y_max
}