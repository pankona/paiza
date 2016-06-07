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

type field struct {
	width, height int
	attr          [][]byte
}

func newField(h, w int) *field {
	var f field
	f.height = h
	f.width = w
	f.attr = make([][]byte, h, h)
	for i := 0; i < h; i++ {
		f.attr[i] = make([]byte, w, w)
	}
	return &f
}

type player struct {
	x, y int
}

func (p *player) String() string {
	return fmt.Sprintf("%d %d", p.x+1, p.y+1)

}

const (
	dirU = "U"
	dirD = "D"
	dirR = "R"
	dirL = "L"
)

func (p *player) move(f field, d string) {
	//fmt.Println("d =", d)
	//fmt.Println("f =", f.width, f.height)
loop:
	for {
		switch d {
		case dirU:
			if p.y-1 >= 0 {
				p.y--
			} else {
				break loop
			}
		case dirD:
			if p.y+1 < f.height {
				p.y++
			} else {
				break loop
			}
		case dirR:
			if p.x+1 < f.width {
				p.x++
			} else {
				break loop
			}
		case dirL:
			if p.x-1 >= 0 {
				p.x--
			} else {
				break loop
			}
		}

		//fmt.Println("p =", p)
		if f.attr[p.y][p.x] == '#' {
			//fmt.Printf("f[%d, %d] is #\n", p.x, p.y)
			continue
		} else {
			break
		}
	}
}

func main() {
	is, e := ReadLineAsIntSlice(" ")
	if e != nil {
		fmt.Println(e)
		return
	}

	h := is[0]
	w := is[1]
	//fmt.Println("h, w =", h, w)

	f := newField(h, w)
	for i := 0; i < h; i++ {
		ss := ReadLineAsStringSlice(" ")
		line := ss[0]
		b := []byte(line)
		for j := 0; j < w; j++ {
			f.attr[i][j] = b[j]
		}
		//fmt.Printf("%s\n", f.attr[i])
	}

	is, e = ReadLineAsIntSlice(" ")
	if e != nil {
		fmt.Println(e)
		return
	}

	p := &player{is[0] - 1, is[1] - 1}
	//fmt.Println("player =", p)

	is, e = ReadLineAsIntSlice(" ")
	if e != nil {
		fmt.Println(e)
		return
	}

	inputNum := is[0]
	//fmt.Println("inputNum =", inputNum)

	for i := 0; i < inputNum; i++ {
		ss := ReadLineAsStringSlice(" ")
		d := ss[0]

		p.move(*f, d)
		//fmt.Println(p)
	}

	fmt.Println(p)
}
