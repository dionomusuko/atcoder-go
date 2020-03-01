package main

import "fmt"

func calc(num int) int {
	var a, b int
	if num < 100 {
		a = num / 10
		b = num % 10
		return a + b
	} else if num >= 100 && num < 1000 {
		a = num / 100
		b = num % 100
		return a + b
	} else {
		a = num / 1000
		b = num % 1000
		return a + b
	}
}

func main() {
	var n, a, b, ans int
	var tmp []int
	fmt.Scan(&n, &a, &b)

	for i := 1; i <= n; i++ {
		if calc(i) >= a && calc(i) <= b {
			tmp = append(tmp, i)
		}
	}
	for i := range tmp {
		ans += tmp[i]
	}
	fmt.Println(tmp)
	fmt.Println(ans)
}
