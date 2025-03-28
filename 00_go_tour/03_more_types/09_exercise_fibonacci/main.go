package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var a, b int = 0, 0
	return func() int {
		if a == 0 && b == 0 {
			a = 1
			return 0
		}
		if b == 0 {
			b = 1
			a = 0
			return 1
		}

		var sum int = a + b
		a = b
		b = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
