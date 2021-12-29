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
	//	"strconv"
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

	Template, Rule, Total := ParseFile(1)
	InsertRule(Template, Rule, 10, Total)
	//fmt.Println(Total)

	Answer1(Total)

	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func InsertRule(template string, rule map[string]string, steps int, total map[string]int) {
	fmt.Println("Template:", template)
	fmt.Println()

	for i := 1; i <= steps; i++ {
		//fmt.Println("Step:",i)
		counter := 0

		for j := 0; j <= (len(template) - counter - 2); j++ {
			pair := template[j+counter : j+2+counter]
			//fmt.Println("Pair:",pair,"J:",j)
			//fmt.Println(template)
			insert := string(rule[pair])

			template = template[:j+1+counter] + insert + template[j+1+counter:]

			total[insert] += 1
			counter += 1

		}
		//fmt.Println(template)
		//fmt.Println("Length:",len(template))
	}

}

func Answer1(total map[string]int) {
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
	fmt.Println("Answer 1:", (most_value - least_value))
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
