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
	fmt.Println("n =", n)
	s := make([][]string, n, n)
	for i := 0; i < n; i++ {
		s[i] = make([]string, 101, 101)
		s[i] = Read(".")
	}

Inspect:
	for i := 0; i < n; i++ {
		v := s[i]
		//fmt.Println("[D] len = ", len(v))
		if len(v) != 4 {
			fmt.Println("False")
			continue
		}

		for j := 0; j < 4; j++ {
			v1, e := strconv.Atoi(v[j])
			//fmt.Println("[D] v, e = ", v1, e)
			if e != nil {
				fmt.Println("False")
				continue Inspect
			} else if v1 < 0 || v1 > 255 {
				fmt.Println("False")
				continue Inspect
			}
		}
		fmt.Println("True")
	}

}
