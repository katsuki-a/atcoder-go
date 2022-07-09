package main

import (
	"bufio"
	"fmt"
	"math"
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
	var i = []int{}
	s := nextStrs(sep)
	for _, v := range s {
		j, _ := strconv.Atoi(v)
		i = append(i, j)
	}
	return i
}

func main() {
	nextInt()
	a := nextInts(" ")
	end := false
	for i := 0; !end; i++ {
		j := math.Pow(2, float64(i+1))
		for _, v := range a {
			if v%int(j) != 0 {
				fmt.Println(i)
				end = true
				break
			}
		}
	}
}
