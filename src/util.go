package main

import "strconv"

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

func sumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
