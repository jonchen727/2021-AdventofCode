package main

import (
	"fmt"
	"os"
	"io/ioutil"
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

const START_CAVE_ID int = 1
const END_CAVE_ID int = 2

func main () {
	start := time.Now()

	lines := FiletoArray("\n",1) 
	//fmt.Println(lines)
	maps := convertLinesToNeighbours(lines)
	fmt.Println(maps)

	fmt.Println(Answer(lines,-1))

	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func getCaveId(caveName string, bigCavePrimeIndex int, smallCavePrimeIndex int) (int, int, int) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541}
	caveId := 0
	
	//fmt.Println(string(caveName[0]))
	
	if caveName == "start" {
		caveId = START_CAVE_ID //if the string is the start cave set it equal to START_CAVE_ID
	} else if caveName == "end" {
		caveId = END_CAVE_ID //if the string is the end cave set it equal to END_CAVE_ID
	} else if caveName[0] > 90 { // if the first ascii number is greater then Z=90 then this is considered lowercase
		caveId = primes[smallCavePrimeIndex] //set the small cave id to the small cave prime index
		smallCavePrimeIndex += 1 //increment the small cave prime index by 1
	} else {
		caveId = primes[bigCavePrimeIndex] // set the big cave id to the big cave prime index
		bigCavePrimeIndex += 1 // increment the big cage prime index by 1 
	}
	return caveId, bigCavePrimeIndex, smallCavePrimeIndex
	
}

func convertLinesToNeighbours(lines []string) map[int][]int {
	bigCavePrimeIndex := 1 //start counting large cave prime index from first index of prime array (2)
	smallCavePrimeIndex := 21 // start counting small cave prime index from 21st index of prime array (79)
	neighbours := map[int][]int{}
	idLookup := map[string]int{}

	for _, line := range lines { //iterate over all lines in input array 
		caveNames := strings.Split(line, "-") // split at - so you have an array of cavenames
		for _, caveName := range caveNames {  // iterate over array of cavenames for each individual cavename

			caveId, ok := idLookup[caveName] //populate the idLookup wih caveName, skipping duplicates
			//fmt.Println("cavename:",caveName,"caveid:",caveId,"ok?",ok)
			//fmt.Println("Before:",idLookup)
			//fmt.Println("Before:",neighbours)

			if !ok {
				caveId, bigCavePrimeIndex, smallCavePrimeIndex = getCaveId(caveName, bigCavePrimeIndex, smallCavePrimeIndex)
				
				idLookup[caveName] = caveId //adds an element with k=caveName, v=caveID to idLookup map, k:v pair 
				//fmt.Println("After:",idLookup)
				neighbours[caveId] = []int{} //creates a empty map int slice for key=caveId

			}

		}



		fromId := idLookup[caveNames[0]] 
		toId := idLookup[caveNames[1]]

		// Appends index 1 to index 0's neighbors 
		if toId != START_CAVE_ID && fromId != END_CAVE_ID { // nothing can go to the start cave/nothing can come from the end cave (if it is not going to start cave, and not coming from end cave then continue)
			neighbours[fromId] = append(neighbours[fromId], toId) //appends the fromId neighbors with the toId 
		
			//fmt.Println("to case",caveNames)
		}

		// Appends index 0 to index 1's neighbors 
		if toId != END_CAVE_ID && fromId != START_CAVE_ID { // nothing can come from the end cave/nothing can go to the start cave (same as above but this is flipped because were adding the inverse case)
			neighbours[toId] = append(neighbours[toId], fromId) //appends the toId neighbors with the fromId
			//fmt.Println("from case",caveNames)
		}
		//fmt.Println(neighbours)
	}
	return neighbours
}

func countPaths(cave int, seen int, allowDoubleVisit int, cache map[int]int, neighbours map[int][]int) int {
	if cave == END_CAVE_ID { //in the event that the path leads to the end cave, add 1 to the count
		fmt.Println("end found")
		fmt.Println("Cave:",cave,"Seen:",seen, "cache:", cache, "neighbors:",neighbours)
		return 1 //return 1 if this leads to end cave
	
	}

	if cave >= 79 { //if cave is a lowercase letter 
		if seen%cave == 0 { //if cave has been visited before, seen%cave means seen*cave has been done before 
			if allowDoubleVisit == -1 { //if allow double visit is -1
				return 0 //return 0 because this path is not valid 
			}
			allowDoubleVisit = -1 //set allowdouble visit to -1 to prevent infinite loop for allowing double visit 
		} else {
			seen *= cave // if the mod of seen does not show the cave has been visited, multiply seen by cave to record a visit
		}
	}
	fmt.Println("Cave:",cave,"Seen:",seen, "cache:", cache, "neighbors:",neighbours)
	total := 0
	for _, neighbour := range neighbours[cave] { //for each neighbor of the cave in question
		//fmt.Println(cave)
		cacheKey := (neighbour+1) * seen * allowDoubleVisit // creates a unique cache key for each neighbor
		
		//cacheKey2 := (neighbour) * seen * allowDoubleVisit
		//fmt.Println(cave,seen, neighbour,cacheKey2,neighbour+1,cacheKey,allowDoubleVisit)
		count, ok := cache[cacheKey] // attempt to find previous record of cache key 
		fmt.Println(cacheKey,count, ok)
		if !ok { // if no pervious record exists 
			count = countPaths(neighbour, seen, allowDoubleVisit, cache, neighbours)
			cache[cacheKey] = count //count either is a 1 or 0 depending on above condition
		}

		total += count
	}
	fmt.Println(cache)
	return total
}
func Answer(lines []string, allowDoubleVisit int) int {
	neighbours := convertLinesToNeighbours(lines)
	cave := START_CAVE_ID
	seen := cave
	cache := map[int]int{}
	
	return countPaths(cave, seen, allowDoubleVisit, cache, neighbours)
}