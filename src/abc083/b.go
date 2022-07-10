package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextStr() string {
	sc.Scan()
	i := sc.Text()
	return i
}

func nextInt() int {
	i, _ := strconv.Atoi(nextStr())
	return i
}

func nextStrs(sep string) []string {
	s := nextStr()
	i := strings.Split(s, sep)
	return i
}

func nextInts(sep string) []int {
	i := []int{}
	s := nextStrs(sep)
	for _, v := range s {
		j, _ := strconv.Atoi(v)
		i = append(i, j)
	}
	return i
}

func intToDigits(i int) []int {
	d := []int{}
	for _, c := range strconv.Itoa(i) {
		d = append(d, int(c)-'0')
	}
	return d
}

func sumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func solve(n int, a int, b int) {
	res := 0
	var sum int
	for i := 1; i < n+1; i++ {
		sum = sumSlice(intToDigits(i))
		if a <= sum && sum <= b {
			res += i
		}
	}
	fmt.Println(res)
}

func main() {
	i := nextInts(" ")
	solve(i[0], i[1], i[2])
}
