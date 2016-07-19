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

	//ip0, _ := strconv.Atoi(ipSS[0])
	rate, _ := strconv.Atoi(ss[0])

	if 0 <= rate && rate <= 30 {
		fmt.Println("sunny")
	} else if 31 <= rate && rate <= 70 {
		fmt.Println("cloudy")
	} else if 71 <= rate {
		fmt.Println("rainy")
	}
}
