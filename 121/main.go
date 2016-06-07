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

	h, _ := strconv.Atoi(ss[0])
	w, _ := strconv.Atoi(ss[1])

	s := make([][]string, h, h)
	for i := 0; i < h; i++ {
		s[i] = make([]string, w, w)
		s[i] = Read(" ")
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			v, _ := strconv.Atoi(s[i][j])
			if v <= 127 {
				fmt.Printf("%d", 0)
			} else {
				fmt.Printf("%d", 1)
			}
			if j+1 != w {
				fmt.Printf(" ")
			}
		}
		if i+1 != h {
			fmt.Printf("\n")
		}
	}

}
