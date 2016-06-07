package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func main() {
	ss := Read(" ")

	itemNum, _ := strconv.Atoi(ss[0])
	userNum, _ := strconv.Atoi(ss[1])
	topNum, _ := strconv.Atoi(ss[2])

	score := Read(" ")

	s := make([][]string, userNum, userNum)
	for i := 0; i < userNum; i++ {
		s[i] = make([]string, 5, 5)
		s[i] = Read(" ")
	}

	userScores := make([]float64, userNum, userNum)
	for i := 0; i < userNum; i++ {
		v := s[i]
		userScores[i] = 0
		for j := 0; j < itemNum; j++ {
			scoreNum, _ := strconv.ParseFloat(score[j], 64)
			score, _ := strconv.ParseFloat(v[j], 64)

			userScores[i] += scoreNum * score
		}
		//fmt.Println("userScore =", userScores[i])
		//fmt.Println("userScore(round) =", round(userScores[i]))
		userScores[i] = round(userScores[i])
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(userScores)))

	for i, score := range userScores {
		if i >= topNum {
			break
		}
		fmt.Println(score)
	}
}
