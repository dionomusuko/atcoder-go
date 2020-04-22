package main

import (
	"fmt"
	"math"
)

func ellipse(radius float64) float64 {
	return radius * 2 * math.Pi
}

func main() {
	var num float64
	fmt.Scan(&num)
	fmt.Println(ellipse(num))
}
