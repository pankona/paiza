package main

import (
	"bufio"
	"errors"
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

// ReadLineAsStringSlice reads stdio and return as slice of string
// input parameter "delim" indicates delimitor for stdio.
// use like: ss := ReadAsStringSlice(" ")
func ReadLineAsStringSlice(delim string) []string {
	ss := strings.Split(readLine(), delim)
	return ss
}

// ReadLineAsIntSlice reads stdio and return as slice of integer
// input parameter "delim" indicates delimitor for stdio.
// use like: ss := ReadAsIntSlice(" ")
func ReadLineAsIntSlice(delim string) ([]int, error) {
	ss := ReadLineAsStringSlice(delim)
	var is []int
	for _, v := range ss {
		iv, e := strconv.Atoi(v)
		if e != nil {
			return nil, errors.New("failed to convert string to integer")
		}
		is = append(is, iv)
	}
	return is, nil
}

func main() {
	totalscore := 0

	is, _ := ReadLineAsIntSlice(" ")
	qNum := is[0]

	for i := 0; i < qNum; i++ {
		answers := ReadLineAsStringSlice(" ")
		//fmt.Println("read:", answers)

		if len(answers[0]) != len(answers[1]) {
			// 0 point
			continue
		}

		if answers[0] == answers[1] {
			// 2 points
			totalscore += 2
			continue
		}

		wrongCharCount := 0
		for j := range answers[0] {
			if answers[0][j] != answers[1][j] {
				wrongCharCount++
			}
		}

		if wrongCharCount == 1 {
			// 1 point
			totalscore++
		}
	}

	fmt.Println(totalscore)
}
