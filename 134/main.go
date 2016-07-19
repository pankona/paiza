package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewReader(os.Stdin)

func readLine() string {
	l, _, _ := sc.ReadLine()
	return string(l)
}

// Read reads stdio and return as slice of string
// delim indicates delimitor for stdio. use like:
// ss := ReadParamsFromStdio(" ")
func Read(delim string) []string {
	ss := strings.Split(readLine(), delim)
	return ss
}

func main() {
	ss := Read(" ")

	name := ss[0]
	gender := ss[1]

	if gender == "M" {
		fmt.Printf("Hi, Mr. %s\n", name)
	} else if gender == "F" {
		fmt.Printf("Hi, Ms. %s\n", name)
	}
}
