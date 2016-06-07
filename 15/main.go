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

func extractIPAddr(s []string) []string {
	return strings.Split(s[0], ".")
}

func extractRangeVal(s []string) (int, int) {
	v1, _ := strconv.Atoi(s[0])
	v2, _ := strconv.Atoi(s[1])
	return v1, v2
}

func main() {
	ipaddrSS := Read(".")
	var ipaddr [2]int
	ipaddr[0], _ = strconv.Atoi(ipaddrSS[0])
	//fmt.Println("ip0 = ", ipaddr[0])
	ipaddr[1], _ = strconv.Atoi(ipaddrSS[1])
	//fmt.Println("ip1 = ", ipaddr[1])

	var ipaddr2Range [2]int
	if strings.HasPrefix(ipaddrSS[2], "[") {
		//fmt.Println("substr = ", ipaddrSS[2][1:len(ipaddrSS[2])-1])
		substr := ipaddrSS[2][1 : len(ipaddrSS[2])-1]
		ipaddr2Range[0], ipaddr2Range[1] = extractRangeVal(strings.Split(substr, "-"))
	} else if ipaddrSS[2] == "*" {
		ipaddr2Range[0] = 0
		ipaddr2Range[1] = 255
	} else {
		ipaddr2Range[0], _ = strconv.Atoi(ipaddrSS[2])
		ipaddr2Range[1], _ = strconv.Atoi(ipaddrSS[2])
	}
	//fmt.Println("ip2 = ", ipaddr2Range)

	var ipaddr3Range [2]int
	if strings.HasPrefix(ipaddrSS[3], "[") {
		//fmt.Println("substr = ", ipaddrSS[2][1:len(ipaddrSS[3])-1])
		substr := ipaddrSS[3][1 : len(ipaddrSS[3])-1]
		ipaddr3Range[0], ipaddr3Range[1] = extractRangeVal(strings.Split(substr, "-"))
	} else if ipaddrSS[3] == "*" {
		ipaddr3Range[0] = 0
		ipaddr3Range[1] = 255
	} else {
		ipaddr3Range[0], _ = strconv.Atoi(ipaddrSS[3])
		ipaddr3Range[1], _ = strconv.Atoi(ipaddrSS[3])
	}
	//fmt.Println("ip3 = ", ipaddr3Range)

	logNumSS := Read(" ")
	logNum, _ := strconv.Atoi(logNumSS[0])

	log := make([][]string, logNum, logNum)
	for i := 0; i < logNum; i++ {
		log[i] = make([]string, 500, 500)
		log[i] = Read(" ")
	}

	for _, v := range log {
		ipSS := extractIPAddr(v)
		ip0, _ := strconv.Atoi(ipSS[0])
		ip1, _ := strconv.Atoi(ipSS[1])
		//ip2, _ := strconv.Atoi(ipSS[2])
		//ip3, _ := strconv.Atoi(ipSS[3])

		//fmt.Println("ip0, ip1 =", ip0, ip1)
		if ipaddr[0] != ip0 || ipaddr[1] != ip1 {
			continue
		}

		ip2, _ := strconv.Atoi(ipSS[2])
		//fmt.Println("ip2 =", ip2)
		//fmt.Println("ip2range =", ipaddr2Range)
		if ipaddr2Range[0] > ip2 || ip2 > ipaddr2Range[1] {
			continue
		}

		ip3, _ := strconv.Atoi(ipSS[3])
		//fmt.Println("ip3 =", ip3)
		//fmt.Println("ip3range =", ipaddr3Range)
		if ipaddr3Range[0] > ip3 || ip3 > ipaddr3Range[1] {
			continue
		}

		fmt.Printf("%s %s %s\n", v[0], v[3][1:len(v[3])], v[6])
	}
}
