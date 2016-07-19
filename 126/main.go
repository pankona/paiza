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
	goodday := 0
	for i := 0; i < 7; i++ {
		ss := Read(" ")
		rate, _ := strconv.Atoi(ss[0])

		if rate <= 30 {
			goodday++
		}
	}

	fmt.Println(goodday)
}
