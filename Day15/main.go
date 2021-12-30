package main

import (
	"fmt"
	//"strconv"
	//	"unicode"
	//"image/color"
	"io/ioutil"
	"os"

	//	"container/heap"
//	"sort"
	"strconv"
	"strings"
	"time"
	//	"github.com/gookit/color"
	//"go.uber.org/zap/internal/color"
)

func ParseFile(arg int) ([][]int) {
	var lines []string
	var ilines [][]int

	if len(os.Args) > arg { // if file argument is provided
		file := os.Args[arg]                //takes 1st arg as file name
		if strings.Contains(file, ".txt") { //checks if file is .txt
			bytes, _ := ioutil.ReadFile(file)      //read file convert to bytes
			input := string(bytes)                 //convert bytes to string
			lines = strings.Split((input), "\n") //convert string to []string with function input as delimiter
			for _,line := range lines {
				var iline []int
				for _,character := range line {
					x,_ := strconv.Atoi(string(character))
					iline = append(iline, x)
				}
				ilines = append(ilines, iline)
			}
		} else { // exit for non text file input
			fmt.Println("Please select a text file") //exits if not .txt file
			os.Exit(69)
		}
	} else { // exit for no argument input
		fmt.Println("Add more files to args")
		os.Exit(420)
	}

	return ilines
}

func main() {
	start := time.Now() //sets current time to start time

	ilines := ParseFile(1)
	

	neighbors := CreateMap(ilines)

	solutions := Solutions{}

	cache := map[string]int{}

	currentmin := 0
	currentminptr := &currentmin

	setCurrentMin(ilines, currentminptr)

	fmt.Println(*currentminptr)

	generateSolutions(ilines,neighbors,0,0,"",0,"",0,solutions, currentminptr,cache)
	fmt.Println(solutions)

	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

type Solutions = map[string]int

type Neighbors = map[string]map[string]int

func generateSolutions (maze [][]int, neighbors Neighbors, x int, y int, path string, risk int, last string, value int, solutions Solutions, currentmin *int, cache map[string]int) {
	newneighbors := Neighbors{}
	
	for k,v := range neighbors {
		newneighbors[k] = v
	}


	//fmt.Println(currentmin)
	//fmt.Println(newneighbors)
	lookupstr := strconv.Itoa(x) + "," + strconv.Itoa(y)
	delete(newneighbors,lookupstr)
	//fmt.Println(neighbors)
	path = path + " " + lookupstr
	fmt.Println(cache)
	risk += value 
	fmt.Println(risk)
	//fmt.Println(cache)
	_,ok := cache[path];
	fmt.Println(ok)
	if !ok {
		if lookupstr == strconv.Itoa(len(maze[0])-1) + "," + strconv.Itoa(len(maze)-1) {
			fmt.Println("endfound new risk",risk)
			*currentmin = risk
			solutions[path] = risk 
	
			//fmt.Println(solutions)
			
		} else if *currentmin > risk  {
			for k, v := range neighbors[lookupstr]{
				//fmt.Println(k,v)
				var newx int
				var newy int
				
				if k == "up" && last != "down" {
					newy = y-1
					last := "up"
					_,ok := newneighbors[strconv.Itoa(x) + "," + strconv.Itoa(newy)];
					if ok {
						generateSolutions(maze,newneighbors,x,newy,path,risk,last,v,solutions,currentmin,cache)
					}
					
				}
				if k == "down" && last != "up"{
					newy = y+1
					last := "down"
					_,ok := newneighbors[strconv.Itoa(x) + "," + strconv.Itoa(newy)];
					if ok{
						generateSolutions(maze,newneighbors,x,newy,path,risk,last,v,solutions,currentmin,cache)
					}
				}
				if k == "left" && last != "right" {
					newx = x-1
					last := "left"
					_,ok := newneighbors[strconv.Itoa(newx) + "," + strconv.Itoa(y)];
					if ok {
						generateSolutions(maze,newneighbors,newx,y,path,risk,last,v,solutions,currentmin,cache)
					}
				}
				if k == "right" && last != "left"{
					newx = x+1
					last := "right"
					_,ok := newneighbors[strconv.Itoa(newx) + "," + strconv.Itoa(y)];
					if ok {
					generateSolutions(maze,newneighbors,newx,y,path,risk,last,v,solutions,currentmin,cache)
					}
				}
			}
		} else {
			fmt.Println("min too high",risk)
			cache[path] = 1
			
		}
	}else {
		fmt.Println("this path is trash")
	}

	

}

func setCurrentMin (maze [][]int, currentmin *int) {
	var sum int
	for index := 0 ; index < len(maze[0]); index++ {
		sum += maze[0][index]
	}
	for line := 0; line < len(maze); line++ {
		sum += maze[line][len(maze[0])-1]
	}
	*currentmin = sum

}
func CreateMap(maze [][]int) Neighbors {
	neighbors := Neighbors{}
	xmax := len(maze[0])
	ymax := len(maze)

	for y ,line := range maze {
		for x, _ := range line {
			
			up := y - 1
			down := y + 1 
			left := x - 1
			right := x + 1
			

			coord := strconv.Itoa(x) + "," + strconv.Itoa(y)
			//fmt.Println(coord)
			neighbors[coord] = map[string]int{}
			if up >= 0 {
				
				v := maze[up][x]
				neighbors[coord]["up"] = v
			}
			if down < ymax {
				v := maze[down][x]
				neighbors[coord]["down"] = v
			}
			if left >= 0 {
				v := maze[y][left]
				neighbors[coord]["left"] = v
			}
			if right < xmax {
				v := maze[y][right]
				neighbors[coord]["right"] = v
			}
		}
	}
	return neighbors
}

