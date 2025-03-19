package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%v is negative, change it!", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
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
	// fmt.Println(i)
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
