package main

import "fmt"

func Vote(ary []int, m int) string {
	var sum int
	for i := 0; i < len(ary); i++ {
		sum += ary[i]
	}
	num := 4 * m
	min := sum / num
	tmp := 1001
	for i := 0; i < len(ary); i++ {
		if tmp > ary[i] {
			tmp = ary[i]
		}
	}
	if tmp < min {
		return "No"
	} else {
		return "Yes"
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	ary := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ary[i])
	}
	fmt.Println(Vote(ary, m))
}
