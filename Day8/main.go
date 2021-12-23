package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

type Signal struct {
	sigpat []string
	output []string
}

type Decoder struct {
	number int
	pat    []string
}

func main() {
	start := time.Now() //sets current time to start time
	lines := (FiletoArray("\n", 1))

	signals := GenerateSignal(lines)

	fmt.Println("Answer 1:", Answer1(signals))
	fmt.Println("Answer 2:", Answer2(signals))

	fmt.Println()
	duration := time.Since(start) //sets duration to time difference since start
	fmt.Println("This Script took:", duration, "to complete!")
}

func Answer2(signals []Signal) int {
	var Answer2 int
	for _, signal := range signals {
		var solution string
		var Decoders [10]Decoder
		Decoders = GenerateDecoder(signal, Decoders)
		for _, output := range signal.output {
			length := len(output)

			switch length {
			case 2:
				solution += "1"
			case 3:
				solution += "7"
			case 4:
				solution += "4"
			case 7:
				solution += "8"
			case 5:
				solution += NumberMapper(Decoders, 2, output)
				solution += NumberMapper(Decoders, 3, output)
				solution += NumberMapper(Decoders, 5, output)
			case 6:
				solution += NumberMapper(Decoders, 6, output)
				solution += NumberMapper(Decoders, 0, output)
				solution += NumberMapper(Decoders, 9, output)
			}

		}
		iSolution, _ := strconv.Atoi(solution)
		Answer2 += iSolution

	}
	return Answer2
}

func NumberMapper(Decoders [10]Decoder, i int, signal string) string {
	ctr := 0
	var output string
	for _, character := range Decoders[i].pat {
		//fmt.Println("im trying",i)
		solution := len(Decoders[i].pat)
		if strings.Contains(signal, character) {
			ctr += 1
		}
		if ctr == solution {
			output = strconv.Itoa(i)
			//fmt.Println("yes",i,"worked")
		}
	}
	return output
}

func GenerateDecoder(signals Signal, Decoders [10]Decoder) [10]Decoder {
	Decoders = DecodeUnique(signals, Decoders)
	Decoders = DecoderRemaining(signals, Decoders)
	return Decoders

}

func DecoderRemaining(signal Signal, Decoders [10]Decoder) [10]Decoder {
	for _, sig := range signal.sigpat {
		length := len(sig)
		var pat []string

		for j := 0; j < len(sig); j++ {
			y := string(sig[j])
			pat = append(pat, y)
		}
		switch length {
		case 6:
			//fmt.Println("might be a 6 or 9 or 0")
			var ctr int
			for _, character := range Decoders[4].pat {

				//fmt.Println(pat)

				if strings.Contains(sig, character) {
					ctr += 1

				}
			}
			//fmt.Println(len(pat))
			if ctr == 4 {
				//fmt.Println(len(pat))
				//fmt.Println("its a 9")
				decode := Decoder{number: 9, pat: pat}
				Decoders[9] = decode

			} else if ctr == 3 {
				//fmt.Println("might be a 6 or 0")
				ctr2 := 0
				for _, character := range Decoders[1].pat {
					if strings.Contains(sig, character) {
						ctr2 += 1
					}
				}
				if ctr2 == 2 {
					//fmt.Println("its a 0")
					decode := Decoder{number: 0, pat: pat}
					Decoders[0] = decode
				} else if ctr2 == 1 {
					//fmt.Println("its a 6")
					decode := Decoder{number: 6, pat: pat}
					Decoders[6] = decode
				}
			}

		case 5:
			//fmt.Println("might be a 3 or a 5 or a 2")
			var ctr int
			for _, character := range Decoders[4].pat {
				if strings.Contains(sig, character) {
					ctr += 1
				}
			}
			//fmt.Println(ctr, sig)
			if ctr == 3 {
				//fmt.Println("might be a 3 or a 5")
				ctr2 := 0
				for _, character := range Decoders[1].pat {
					if strings.Contains(sig, character) {
						ctr2 += 1
					}
				}
				if ctr2 == 1 {
					//fmt.Println("its a 5")
					decode := Decoder{number: 5, pat: pat}
					Decoders[5] = decode
				} else if ctr2 == 2 {
					//fmt.Println("its a 3")
					decode := Decoder{number: 3, pat: pat}
					Decoders[3] = decode
				}

			} else if ctr == 2 {
				//fmt.Println("FUCK", ctr)
				//fmt.Println("its a 2")
				decode := Decoder{number: 2, pat: pat}
				Decoders[2] = decode
			}
		}

	}
	return Decoders
}

func DecodeUnique(signals Signal, Decoders [10]Decoder) [10]Decoder {
	for i, sig := range signals.sigpat {
		length := len(sig)
		var pat []string
		switch length {
		case 2:
			//fmt.Println("its a 1")
			for _, character := range signals.sigpat[i] {
				x := string(character)
				//fmt.Println(x)
				pat = append(pat, x)
			}
			decode := Decoder{number: 1, pat: pat}
			Decoders[1] = decode

		case 3:
			//fmt.Println("its a 7")
			for _, character := range signals.sigpat[i] {
				x := string(character)
				//fmt.Println(x)
				pat = append(pat, x)
			}
			decode := Decoder{number: 7, pat: pat}
			Decoders[7] = decode

		case 7:
			//fmt.Println("its an 8")
			for _, character := range signals.sigpat[i] {
				x := string(character)
				//fmt.Println(x)
				pat = append(pat, x)
			}
			decode := Decoder{number: 8, pat: pat}
			Decoders[8] = decode

		case 4:
			//fmt.Println("its an 4")
			for _, character := range signals.sigpat[i] {
				x := string(character)
				//fmt.Println(x)
				pat = append(pat, x)
			}
			decode := Decoder{number: 4, pat: pat}
			Decoders[4] = decode

		}
	}
	return Decoders
}

func Answer1(signals []Signal) int {
	var ctr int
	for _, signal := range signals {
		for _, line := range signal.output {

			seg := len(line)

			if seg == 7 || seg == 3 || seg == 2 || seg == 4 {
				ctr += 1
				//fmt.Println(line, seg, "*")
			} else {
				//fmt.Println(line, seg)
			}

		}
	}
	return ctr
}

func GenerateSignal(lines []string) []Signal {
	var lines2 [][]string
	for _, line := range lines {
		split := strings.Split((line), "|")
		lines2 = append(lines2, split)
	}

	var Signals []Signal

	for _, line := range lines2 {
		sigpat := strings.Fields(line[0])
		output := strings.Fields(line[1])

		signal := Signal{sigpat: sigpat, output: output}
		Signals = append(Signals, signal)

	}
	return Signals
}
