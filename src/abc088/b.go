package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func solve(a []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	alice := 0
	bob := 0
	for i, card := range a {
		if i%2 == 0 {
			alice += card
		} else {
			bob += card
		}
	}
	fmt.Println(alice - bob)
}

func main() {
	nextInt()
	a := nextInts(" ")
	solve(a)
}
