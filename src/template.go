package main

import (
	"bufio"
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

func nextStrs(sep string) []string {
	s := nextStr()
	i := strings.Split(s, sep)
	return i
}

func nextStrLines(n int) []string {
	s := []string{}
	for i := 0; i < n; i++ {
		s = append(s, nextStr())
	}
	return s
}

func nextInt() int {
	i, _ := strconv.Atoi(nextStr())
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

func nextIntLines(n int) []int {
	s := []int{}
	for i := 0; i < n; i++ {
		s = append(s, nextInt())
	}
	return s
}

func solve() {}

func main() {}
