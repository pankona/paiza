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

func getCardVal(cards [][]int, h, w int) int {
	//fmt.Println("h, w =", h, w)
	return cards[h-1][w-1]
}

func main() {
	ss := Read(" ")

	h, _ := strconv.Atoi(ss[0])
	w, _ := strconv.Atoi(ss[1])
	playerNum, _ := strconv.Atoi(ss[2])

	//fmt.Println("h, w, playerNum = ", h, w, playerNum)

	cardsSS := make([][]string, h, h)
	cards := make([][]int, h, h)
	for i := 0; i < h; i++ {
		cardsSS[i] = make([]string, w, w)
		cardsSS[i] = Read(" ")

		cards[i] = make([]int, w, w)
		for j := 0; j < w; j++ {
			v, _ := strconv.Atoi(cardsSS[i][j])
			cards[i][j] = v
		}
		//fmt.Println(cards[i])
	}

	recordLenSS := Read(" ")
	recordLen, _ := strconv.Atoi(recordLenSS[0])
	//fmt.Println("recordLen =", recordLen)

	inputSS := make([][]string, recordLen, recordLen)
	input := make([][]int, recordLen, recordLen)
	playerScore := make([]int, playerNum, playerNum)
	currentPlayer := 0
	for i := 0; i < recordLen; i++ {
		inputSS[i] = make([]string, 4, 4)
		inputSS[i] = Read(" ")
		//fmt.Println("inputSS[i] =", inputSS[i])

		input[i] = make([]int, 4, 4)
		for j := 0; j < 4; j++ {
			v, _ := strconv.Atoi(inputSS[i][j])
			input[i][j] = v
		}

		//fmt.Println(input[i])

		//fmt.Println("cards =", cards)
		//fmt.Println("input =", input)

		c1 := getCardVal(cards, input[i][0], input[i][1])
		c2 := getCardVal(cards, input[i][2], input[i][3])
		//fmt.Println("c1, c2 =", c1, c2)
		//fmt.Println("currentPlayer =", currentPlayer)
		if c1 == c2 {
			playerScore[currentPlayer] += 2
		} else {
			currentPlayer++
			currentPlayer %= playerNum
		}
	}

	for i := 0; i < playerNum; i++ {
		fmt.Println(playerScore[i])
	}
}
