package main

import (
	"os"
	"fmt"
)

func main () {
	somestring := string(os.Args[1])
	fmt.Println(somestring)

}