package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	var t T

	i = &T{"Hello"}
	describe(i)
	i.M()
	t = T{"Hello"}
	describe2(t)
	t.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe2(i T) {
	fmt.Printf("(%v, %T)\n", i, i)
}
