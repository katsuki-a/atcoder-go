package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func intToDigits(i int) []int {
	d := []int{}
	for _, c := range strconv.Itoa(i) {
		d = append(d, int(c)-'0')
	}
	return d
}

func reverseSlice(slice []int) []int {
	reversed := []int{}
	for _, v := range slice {
		reversed = append([]int{v}, reversed...)
	}
	return reversed
}

func solve(a int, b int) {
	res := 0
	var digits []int
	for i := a; i <= b; i++ {
		digits = intToDigits(i)
		reversed := reverseSlice(digits)
		if reflect.DeepEqual(digits, reversed) {
			res += 1
		}
	}
	fmt.Println(res)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	solve(a, b)
}
