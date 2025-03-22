package main

import (
	"errors"
	"fmt"
	"net/http"
)

const port = ":8080"

// Home is the homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	var sum int = sumVals(12, 5)
	fmt.Fprintf(w, "This is Home and 12 + 5 is %d", sum)
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	var sum int = sumVals(4, 6)
	fmt.Fprintf(w, "This is About and 4 + 6 is %d", sum)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	result, err := divide(10.0, 2.0)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, result)
}

func BadDivide(w http.ResponseWriter, r *http.Request) {
	result, err := divide(10.0, 0.0)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, result)
}

// divide divides two floats
func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("can not divide by 0")
	}
	return a / b, nil
}

// sumVals sums to integers
func sumVals(a, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	http.HandleFunc("/baddivide", BadDivide)

	fmt.Println("Server running on", port)
	http.ListenAndServe(port, nil)
}
