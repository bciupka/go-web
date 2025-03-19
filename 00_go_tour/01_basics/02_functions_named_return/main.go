package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func mark(x, y string) (a, b string) {
	a = x + "m"
	b = y + "m"
	return
}

func main() {
	fmt.Println(split(17))
	fmt.Println(mark("foo", "bar"))
}
