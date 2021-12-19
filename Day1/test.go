package main

import (
	"os"
	"fmt"
)

func main () {
	somestring := os.Args[1]
	fmt.Println(somestring)
}