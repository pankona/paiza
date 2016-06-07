package main

import (
	"bufio"
	"errors"
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

type point struct {
	x, y     int
	value    int
	distance float64
}

func (p point) String() string {
	return fmt.Sprintf("[%d, %d], value = %d, distance = %f", p.x, p.y, p.value, p.distance)
}

func (p point) calcDistance(p2 point) float64 {
	return math.Sqrt(float64((p.x-p2.x)*(p.x-p2.x) + (p.y-p2.y)*(p.y-p2.y)))
}

// sort for point struct
type points []point

func (ps points) Len() int {
	return len(ps)
}

func (ps points) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps points) Less(i, j int) bool {
	return ps[i].distance < ps[j].distance
}

func round(f float64) int {
	return int(math.Floor(f + .5))
}

func main() {
	is, _ := ReadLineAsIntSlice(" ")
	a := point{is[0], is[1], 0, 0}
	//fmt.Println("a =", a)

	is, _ = ReadLineAsIntSlice(" ")
	k := is[0]
	//fmt.Println("k = ", k)

	is, _ = ReadLineAsIntSlice(" ")
	N := is[0]

	var points points
	points = make([]point, N, N)
	for i := 0; i < N; i++ {
		is, _ := ReadLineAsIntSlice(" ")
		points[i].x = is[0]
		points[i].y = is[1]
		points[i].value = is[2]
		points[i].distance = a.calcDistance(points[i])
		//fmt.Println("point =", points[i])
	}

	sort.Sort(points)

	var value int
	for i := 0; i < k; i++ {
		value += points[i].value
	}
	//fmt.Println("value =", value)
	//fmt.Println("k =", k)

	fmt.Println(round(float64(value) / float64(k)))
}
