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

func blackToWhite(coins string, boardLen int, result []byte) []byte {
	bIndexFrom := 0
	bIndexTo := 0
	for i := 0; i < boardLen-1; i++ {
		if coins[i:i+2] == "wb" {
			//fmt.Println("wb found")
			bIndexFrom = i + 1
		}
		if bIndexFrom != 0 && coins[i:i+2] == "bw" {
			//fmt.Println("bw found")
			bIndexTo = i + 1
			//fmt.Println("from, to =", bIndexFrom, bIndexTo)
			for j := bIndexFrom; j < bIndexTo; j++ {
				//fmt.Println("@@@ black to white")
				result[j] = 'w'
			}
			bIndexFrom = 0
			bIndexTo = 0
		}
	}

	return result
}

func whiteToBlack(coins string, boardLen int, result []byte) []byte {
	bIndexFrom := 0
	bIndexTo := 0
	for i := 0; i < boardLen-1; i++ {
		if coins[i:i+2] == "bw" {
			//fmt.Println("bw found")
			bIndexFrom = i + 1
		}
		if bIndexFrom != 0 && coins[i:i+2] == "wb" {
			//fmt.Println("wb found")
			bIndexTo = i + 1
			//fmt.Println("from, to =", bIndexFrom, bIndexTo)
			for j := bIndexFrom; j < bIndexTo; j++ {
				//fmt.Println("@@@ white to black")
				result[j] = 'b'
			}
			bIndexFrom = 0
			bIndexTo = 0
		}
	}

	return result
}

func merge(result1 []byte, result2 []byte, boardLen int) []byte {
	result := make([]byte, boardLen, boardLen)
	for i := 0; i < boardLen; i++ {
		if result1[i] == 'b' {
			result[i] = 'b'
		} else if result2[i] == 'w' {
			result[i] = 'w'
		} else {
			result[i] = 'u'
		}
	}
	return result
}

func restore(coins []byte, result []byte, boardLen int) []byte {
	ret := make([]byte, boardLen, boardLen)
	for i := 0; i < boardLen; i++ {
		if result[i] == 'u' {
			ret[i] = coins[i]
		} else {
			ret[i] = result[i]
		}
	}
	return ret
}

func countB(result []byte) int {
	ret := 0
	for i := 0; i < len(result); i++ {
		if result[i] == 'b' {
			ret++
		}
	}
	return ret
}

func main() {
	boardLenIS, e := ReadLineAsIntSlice(" ")
	if e != nil {
		//fmt.Println(e)
		return
	}
	boardLen := boardLenIS[0]

	coinsSS := ReadLineAsStringSlice(" ")
	coins := coinsSS[0]

	result1 := make([]byte, boardLen, boardLen)
	result2 := make([]byte, boardLen, boardLen)

	//fmt.Println("boardLen =", boardLen)
	//fmt.Println("coins =", coins)

	result := []byte(coins)

	//fmt.Println("result(0) =", string(result1[:boardLen]))
	//fmt.Println("result(0) =", string(result2[:boardLen]))

	for {

		// initialize
		for i := 0; i < boardLen; i++ {
			result1[i] = 'u' // unknown
			result2[i] = 'u' // unknown
		}

		result1 = whiteToBlack(string(result[:boardLen]), boardLen, result1)
		//fmt.Println("result(1) =", string(result1[:boardLen]))

		result2 = blackToWhite(string(result[:boardLen]), boardLen, result2)
		//fmt.Println("result(2) =", string(result2[:boardLen]))

		result3 := merge(result1, result2, boardLen)
		//fmt.Println("result(3) =", string(result3[:boardLen]))

		result3 = restore(result, result3, boardLen)
		//fmt.Println("restored result(3) =", string(result3[:boardLen]))

		//fmt.Println("result, result3 =", string(result[:boardLen]), string(result3[:boardLen]))
		if string(result[:boardLen]) == string(result3[:boardLen]) {
			//fmt.Println("process finished")
			break
		}
		result = result3
		//fmt.Println("result, result3 =", string(result[:boardLen]), string(result3[:boardLen]))
	}
	//fmt.Println("result =", string(result[:boardLen]))

	bNum := countB(result)
	fmt.Println(bNum)

}
