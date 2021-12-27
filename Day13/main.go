package main

import (
	"fmt"
	//	"unicode"
	//"image/color"
	"io/ioutil"
	"os"

	//	"sort"
	"strconv"
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

func main() {
	start := time.Now() //sets current time to start time
	lines := (FiletoArray("\n\n", 1))
	//ilines := StringtoInt(lines)
	//fmt.Println(lines[1])

	dots := CreateDotArray(lines)
	//fmt.Println(dots)

	xmax, ymax := findMax(dots)
	//fmt.Println(xmax,ymax)

	grid := (createGrid(dots, xmax, ymax))

	//fmt.Println(grid)

	folds := ExtractFolds(lines)

	fmt.Println("Answer 1:", Answer1(Folder(grid, folds, 1)))

	Answer2(grid, folds)
	//fmt.Println(xFold(yFold(grid,7),5))

	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func Answer1(grid [][]int) int {
	var count int
	for _, lines := range grid {
		for _, character := range lines {
			if character > 0 {
				count += 1
			}
		}
	}
	return count
}

func Answer2(grid [][]int, folds [][]int) {
	numfold := len(folds)
	newGrid := Folder(grid, folds, numfold)

	for _, line := range newGrid {
		for _, character := range line {
			if character > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func Folder(grid [][]int, folds [][]int, numfold int) [][]int {
	for i := 0; i < numfold; i++ {
		//fmt.Println("Folder run",i)
		if folds[i][0] == 1 {
			grid = yFold(grid, folds[i][1])
			//fmt.Println("yfold")
		}
		if folds[i][0] == 0 {
			grid = xFold(grid, folds[i][1])
		}
	}
	return grid
}

func yFold(grid [][]int, position int) [][]int {
	xmax := len(grid[0])
	ymax := len(grid)
	yeven := ymax % 2
	//fmt.Println(xmax,yeven)
	//fmt.Println(xmax,(ymax-yeven)/2)
	folded := make([][]int, ymax/2)
	for i := range folded {
		folded[i] = make([]int, xmax)
	}

	for y, line := range grid {
		for x, value := range line {
			if y < ymax/2 && value > 0 {
				folded[y][x] += grid[y][x]
				//fmt.Println("x:",x,"y:",y,"value:",grid[y][x])
				//fmt.Println("xnew:",x,"ynew:",y)
			} else if y > ymax/2 && value > 0 {
				folded[(ymax-yeven)-y][x] += grid[y][x]
				//fmt.Println("x:",x,"y:",y,"value:",grid[y][x])
				//fmt.Println("xnew:",x,"ynew:",(ymax-1)-y)

			}

		}
	}
	return folded

}
func xFold(grid [][]int, position int) [][]int {
	xmax := len(grid[0])
	ymax := len(grid)
	xeven := xmax % 2
	//fmt.Println(xmax,ymax)
	//fmt.Println(xmax/2,ymax)
	folded := make([][]int, ymax)
	for i := range folded {
		folded[i] = make([]int, ((xmax - xeven) / 2))
	}

	for y, line := range grid {

		for x, value := range line {
			//fmt.Println(xmax)
			//fmt.Println("x:",x,"y:",y,"value:",grid[y][x])
			//fmt.Println("xnew:",x,"ynew:",y)
			if x < xmax/2 && value > 0 {
				folded[y][x] += grid[y][x]
				//fmt.Println("x:",x,"y:",y,"value:",grid[y][x])
				//fmt.Println("xnew:",x,"ynew:",y)
			} else if x > xmax/2 && value > 0 {
				folded[y][(xmax-xeven)-x] += grid[y][x]
				//fmt.Println("x:",x,"y:",y,"value:",grid[y][x])
				//fmt.Println("xnew:",(xmax-1)-x,"ynew:",y)

			}

		}
	}
	return folded

}

func ExtractFolds(lines []string) [][]int {
	var Folds [][]int

	for _, line := range strings.Split(lines[1], "\n") {
		folds := line[11:]
		fold := strings.Split(folds, "=")

		direction := fold[0]
		value, _ := strconv.Atoi(fold[1])

		var iDirection int

		if direction == "x" {
			iDirection = 0
		} else if direction == "y" {
			iDirection = 1
		}

		Folds = append(Folds, []int{iDirection, value})
	}
	return Folds
}

func CreateDotArray(lines []string) [][]int {
	var dots [][]int

	for _, dot := range strings.Split(lines[0], "\n") {
		coords := strings.Split(dot, ",")
		var icoords []int
		for _, coord := range coords {
			x, _ := strconv.Atoi(coord)
			icoords = append(icoords, x)

		}
		//fmt.Println(icoords)
		dots = append(dots, icoords)

	}
	return dots
}

func createGrid(dots [][]int, xmax int, ymax int) [][]int {
	grid := make([][]int, ymax)
	for i := range grid {
		grid[i] = make([]int, xmax)
	}

	for _, dot := range dots {
		x := dot[0]
		y := dot[1]

		grid[y][x] += 1
	}

	return grid
}

func findMax(dots [][]int) (int, int) {
	xmax := dots[0][0]
	ymax := dots[0][1]

	for _, value := range dots {
		if value[0] > xmax {
			xmax = value[0]
		}
		if value[1] > ymax {
			ymax = value[1]
		}
	}

	return xmax + 1, ymax + 1
}
