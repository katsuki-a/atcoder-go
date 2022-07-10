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

func main() {
	a := nextInt()
	b := nextInt()
	c := nextInt()
	x := nextInt()
	res := 0
	for i := 0; i < a+1; i++ {
		for j := 0; j < b+1; j++ {
			for k := 0; k < c+1; k++ {
				if i*500+j*100+k*50 == x {
					res += 1
				}
			}
		}
	}
	fmt.Println(res)
}
