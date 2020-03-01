package main

import (
	"fmt"
)

func fizz(ary []int) int {
	var count int
	for i := range ary {
		if ary[i]%2 == 0 {
			count += 1
		} else {
			break
		}
	}
	return count
}

func division(ary []int) []int {
	var tmp []int
	for i := range ary {
		tmp = append(tmp, ary[i]/2)
	}
	return tmp
}

func main() {
	var n int
	count := 0
	fmt.Scan(&n)
	// 1, 2, 3 ... の入力法
	ary := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ary[i])
	}
	for {
		if len(ary) == fizz(ary) {
			ary = division(ary)
			count += 1
		} else {
			break
		}
	}
	fmt.Println(count)
}
