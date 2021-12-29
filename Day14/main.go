package main

import (
	"fmt"
	//"strconv"
	//	"unicode"
	//"image/color"
	"io/ioutil"
	"os"

	//	"container/heap"
	"sort"
	"strconv"
	"strings"
	"time"
	//	"github.com/gookit/color"
	//"go.uber.org/zap/internal/color"
)

func ParseFile(arg int) (string, map[string]string, map[string]int) {
	var lines []string
	lookup := map[string]string{}
	total := map[string]int{}

	if len(os.Args) > arg { // if file argument is provided
		file := os.Args[arg]                //takes 1st arg as file name
		if strings.Contains(file, ".txt") { //checks if file is .txt
			bytes, _ := ioutil.ReadFile(file)      //read file convert to bytes
			input := string(bytes)                 //convert bytes to string
			lines = strings.Split((input), "\n\n") //convert string to []string with function input as delimiter
		} else { // exit for non text file input
			fmt.Println("Please select a text file") //exits if not .txt file
			os.Exit(69)
		}
	} else { // exit for no argument input
		fmt.Println("Add more files to args")
		os.Exit(420)
	}

	startstr := (lines[0])
	imtired := strings.Split(lines[1], "\n")

	for _, line := range imtired {
		imtired = strings.Split(line, " -> ")

		lookup[string(imtired[0])] = string(imtired[1])
	}

	for _, character := range startstr {
		total[string(character)] += 1

	}

	return startstr, lookup, total
}

func main() {
	start := time.Now() //sets current time to start time

	cache := Cache{}

	Template, Rule, Total := ParseFile(1)

	fmt.Println("Answer 1")
	Total1 := addtoMap(InsertRule(Template, Rule, 10, cache), Total)
	mostMinusleast(Total1)

	fmt.Println()
	fmt.Println("Answer 2")

	Total2 := addtoMap(InsertRule(Template, Rule, 40, cache), Total)
	mostMinusleast(Total2)

	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

type Cache = map[string]map[string]int
type Total = map[string]int

func InsertRule(template string, rule map[string]string, steps int, cache Cache) map[string]int {

	if steps == 0 {
		return Total{}
	}

	cacheKey := template + strconv.Itoa(steps)
	if total2, ok := cache[cacheKey]; ok {
		return total2
	}

	steps -= 1

	total2 := Total{}

	for i := 0; i < len(template)-1; i++ {
		front := string(template[i])
		back := string(template[i+1])

		if insert, ok := rule[front+back]; ok {
			total2[insert] += 1
			total2 = addtoMap(total2, InsertRule(front+insert+back, rule, steps, cache))
		}
	}
	cache[cacheKey] = total2

	return total2

}

func addtoMap(total1 map[string]int, total2 map[string]int) map[string]int {

	for k, v := range total2 {
		total1[k] += v
	}
	return total1
}

func mostMinusleast(total map[string]int) {
	p := make(PairList, len(total))
	i := 0
	for k, v := range total {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	most_index := len(p) - 1
	most_value := p[most_index].Value
	least_value := p[0].Value
	fmt.Println("Most:", p[most_index].Key, "at", most_value)
	fmt.Println("Least:", p[0].Key, "at", least_value)
	fmt.Println()
	fmt.Println("Most-Least:", (most_value - least_value))
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
