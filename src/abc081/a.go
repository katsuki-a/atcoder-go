package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextStr() string {
	sc.Scan()
	i := sc.Text()
	return i
}

func main() {
	i := nextStr()
	a := strings.Split(i, "")
	res := 0
	for _, v := range a {
		b, _ := strconv.Atoi(v)
		res += b
	}
	fmt.Println(res)
}
