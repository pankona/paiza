package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
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

// Round rounds off to the nearest whole number
func Round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}

func main() {
	is, e := ReadLineAsIntSlice(" ")
	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Println(is)
}
