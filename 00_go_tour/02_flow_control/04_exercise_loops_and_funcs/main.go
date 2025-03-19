package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		// fmt.Println(z)
	}
	return z
}

func Sqrt2(x float64) float64 {
	var z float64 = 1
	var tmp float64
	var diff float64 = 1
	i := 0
	for ; diff > 1e-8; i++ {
		tmp = z
		z -= (z*z - x) / (2 * z)
		diff = z - tmp
		if diff < 0 {
			diff = -diff
		}
	}
	fmt.Println(i)
	return z
}

func main() {
	x := 10.0
	fmt.Println(math.Sqrt(x))
	fmt.Println(Sqrt(x))
	fmt.Println(Sqrt2(x))
}
