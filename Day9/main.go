package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
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
	lines := (FiletoArray("\n", 1))
	//ilines := StringtoInt(lines)
	var LowPoints []LowPoint
	var sum int
	sum, LowPoints = Answer1(lines, LowPoints)
	fmt.Println("Answer 1:", sum)

	fmt.Println("Answer 2:", Answer2(LowPoints))

	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

type LowPoint struct {
	point      []int
	valleysize int
}

func Answer1(lines []string, LowPoints []LowPoint) (int, []LowPoint) {
	var sum int
	var counter int

	for l, line := range lines {

		for i, index := range line {
			var right string
			var left string
			var top string
			var bottom string
			var neighbors float64
			var lower float64
			var coords [][]int
			var lower2 int

			x := string(index)
			ix, _ := strconv.Atoi(x)
			if i < len(string(line))-1 {
				right = string(lines[l][i+1])
				iright, _ := strconv.Atoi(right)

				neighbors += 1
				if iright > ix {
					lower += 1
					if iright < 9 {
						var coord []int
						lower2 += 1
						coord = append(coord, i+1, l, iright)
						coords = append(coords, coord)
					}
				}

			}
			if i > 0 {
				left = string(lines[l][i-1])
				ileft, _ := strconv.Atoi(left)
				neighbors += 1
				if ileft > ix {
					lower += 1
					if ileft < 9 {
						var coord []int
						lower2 += 1
						coord = append(coord, i-1, l, ileft)
						coords = append(coords, coord)
					}
				}

			}

			if l < len(lines)-1 {
				bottom = string(lines[l+1][i])
				ibottom, _ := strconv.Atoi(bottom)
				neighbors += 1
				if ibottom > ix {
					lower += 1
					if ibottom < 9 {
						var coord []int
						lower2 += 1
						coord = append(coord, i, l+1, ibottom)
						coords = append(coords, coord)
					}
				}

			}

			if l > 0 {
				top = string(lines[l-1][i])
				itop, _ := strconv.Atoi(top)
				neighbors += 1
				if itop > ix {
					lower += 1
					if itop < 9 {
						var coord []int
						lower2 += 1
						coord = append(coord, i, l-1, itop)
						coords = append(coords, coord)
					}
				}

			}

			if neighbors == lower {
				var xy []int
				xy = append(xy, i, l)
				lowpoint := LowPoint{point: xy}
				LowPoints = append(LowPoints, lowpoint)

				ValleyInspector(coords, lines, LowPoints, counter)

				sum += (ix + 1)
				counter += 1

			}
		}

	}
	return sum, LowPoints
}

func ValleyInspector(coords [][]int, lines []string, LowPoints []LowPoint, counter int) {
	var iteration int

	coordsfinal := coords

	for {

		//fmt.Println("Iteration:", iteration)
		//fmt.Println("Input:", coords)

		coordsprev := LowerFinder(coords, lines, LowPoints, counter)
		//fmt.Println(coordsprev)
		//fmt.Println("Len CoordsPrev:",len(coordsprev))
		coords = removeDuplicateNestedInt(coordsprev)
		//fmt.Println("Len Coords:",len(coords))
		//fmt.Println("Output:", coords)
		for _, coord := range coords {
			coordsfinal = append(coordsfinal, coord)
		}
		coordsfinal = removeDuplicateNestedInt(coordsfinal)

		iteration += 1
		if len(coords) == 0 {
			//fmt.Println("Final:", coordsfinal, len(coordsfinal))
			LowPoints[counter].valleysize = len(coordsfinal) + 1
			break
		}

	}
}

func LowerFinder(coords [][]int, lines []string, LowPoints []LowPoint, counter int) [][]int {
	var coords2 [][]int
	for _, coord := range coords {
		i := coord[0]
		l := coord[1]
		ix := coord[2]
		var right string
		var left string
		var top string
		var bottom string
		var neighbors float64
		var lower float64
		line := lines[0]

		lower = 0

		if i < len(string(line))-1 {
			right = string(lines[l][i+1])
			iright, _ := strconv.Atoi(right)

			if iright != 9 && iright > ix {
				lower += 1
				var coord2 []int
				coord2 = append(coord2, i+1, l, iright)
				coords2 = append(coords2, coord2)
			}

		}
		if i > 0 {
			left = string(lines[l][i-1])
			ileft, _ := strconv.Atoi(left)
			neighbors += 1
			if ileft != 9 && ileft > ix {
				lower += 1
				var coord2 []int
				coord2 = append(coord2, i-1, l, ileft)
				coords2 = append(coords2, coord2)
			}

		}

		if l < len(lines)-1 {
			bottom = string(lines[l+1][i])
			ibottom, _ := strconv.Atoi(bottom)
			neighbors += 1
			if ibottom != 9 && ibottom > ix {
				lower += 1
				var coord2 []int
				coord2 = append(coord2, i, l+1, ibottom)
				coords2 = append(coords2, coord2)
			}

		}

		if l > 0 {
			top = string(lines[l-1][i])
			itop, _ := strconv.Atoi(top)
			neighbors += 1
			if itop != 9 && itop > ix {
				lower += 1
				var coord2 []int
				coord2 = append(coord2, i, l-1, itop)
				coords2 = append(coords2, coord2)
			}

		}

	}
	return coords2
}

func removeDuplicateNestedInt(coords [][]int) [][]int {
	var strSlice []string
	newcords := [][]int{}
	for _, coord := range coords {
		var newstr string
		for _, nums := range coord {
			snums := strconv.Itoa(nums)
			newstr += snums

		}
		strSlice = append(strSlice, newstr)
	}

	allKeys := make(map[string]bool)
	list := []string{}
	for i, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
			newcords = append(newcords, coords[i])

			//fmt.Println(i,value,list)
		} else {
			//fmt.Println(i,value,list)
		}

	}
	//fmt.Println(newcords)
	return newcords
}

func Answer2(LowPoints []LowPoint) int {
	product := 1
	sort.Slice(LowPoints[:], func(i, j int) bool {
		return LowPoints[i].valleysize > LowPoints[j].valleysize
	})

	for i := 0; i < 3; i++ {
		product *= LowPoints[i].valleysize
	}
	return product
}
