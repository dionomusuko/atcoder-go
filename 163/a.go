package main

import (
	"fmt"
	"math"
)

//面積を求める関数
func ellipse(radius float64) float64 {
	return radius * 2 * math.Pi
}

func main() {
	var num float64
	fmt.Scan(&num)
	fmt.Println(ellipse(num))
}
