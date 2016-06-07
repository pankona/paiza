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

	memberNum, _ := strconv.Atoi(ss[0])
	liveNum, _ := strconv.Atoi(ss[1])

	s := make([][]string, liveNum, liveNum)
	for i := 0; i < liveNum; i++ {
		s[i] = make([]string, memberNum, memberNum)
		s[i] = Read(" ")
		//fmt.Println(s[i])
	}

	scores := make([]int, liveNum, liveNum)
	for i := 0; i < liveNum; i++ {
		v := s[i]
		scores[i] = 0
		tmpscore := 0
		for j := 0; j < memberNum; j++ {
			//fmt.Println("v =", v)
			tmpscore, _ = strconv.Atoi(v[j])
			scores[i] += tmpscore
		}
		if scores[i] < 0 {
			scores[i] = 0
		}
		//fmt.Println("scores[i] =", scores[i])
	}

	var result int
	for _, v := range scores {
		result += v
	}
	fmt.Println(result)
}
