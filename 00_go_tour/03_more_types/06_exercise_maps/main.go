package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var fields []string = strings.Fields(s)
	var m map[string]int = make(map[string]int)
	for _, e := range fields {
		m[e] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
