package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	describe(42)
	describe("hello")
	fmt.Printf("(%v, %T)\n", 42, 42)
	fmt.Printf("(%v, %T)\n", "hello", "hello")
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
