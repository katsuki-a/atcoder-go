package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)

func nextStr() string {
	sc.Scan()
	i := sc.Text()
	return i
}

func nextStrLines(n int) []string {
	s := []string{}
	for i := 0; i < n; i++ {
		s = append(s, nextStr())
	}
	return s
}

func solve(s []string) {
	sort.Strings(s)
	for _, v := range s {
		fmt.Print(v)
	}
}

func main() {
	var n, l int
	fmt.Scan(&n, &l)
	s := nextStrLines(n)
	solve(s)
}
