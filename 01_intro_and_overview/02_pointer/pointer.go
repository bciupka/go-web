package main

import "fmt"

func main() {
	var car string = "car"
	maybeChanged := maybeChage(car)
	fmt.Println(maybeChanged)
	fmt.Println(car)

	changed := change(&car)
	fmt.Println(changed)
	fmt.Println(car)
}

func maybeChage(s string) string {
	s = "truck"
	return s
}

func change(s *string) string {
	*s = "bike"
	return *s
}