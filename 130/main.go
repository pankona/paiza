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
	a, _ := strconv.Atoi(ss[0])
	b, _ := strconv.Atoi(ss[1])

	ss = Read(" ")
	c, _ := strconv.Atoi(ss[0])
	d, _ := strconv.Atoi(ss[1])

	fmt.Println(a*d - b*c)
}
