package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func nextIntLines(n int) []int {
	s := []int{}
	for i := 0; i < n; i++ {
		s = append(s, nextInt())
	}
	return s
}

func solve(d []int) {
	set := make(map[int]struct{})
	for _, v := range d {
		set[v] = struct{}{}
	}
	fmt.Println(len(set))
}

func main() {
	n := nextInt()
	d := nextIntLines(n)
	solve(d)
}
