package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	n, _ := strconv.Atoi(ss[0])
	g := "Unknown"
	switch n {
	case 1:
		g = "E"
	case 2:
		g = "D"
	case 3:
		g = "C"
	case 4:
		g = "B"
	case 5:
		g = "A"
	default:
		fmt.Println("Unknown value")
		os.Exit(1)
	}

	fmt.Println(g)
}
