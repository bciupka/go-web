package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintln(w, "Hello world")
		if err != nil {
			fmt.Println("Error there", err)
		}
		fmt.Printf("%d bytes written\n", n)
	})

	http.ListenAndServe(":8080", nil)
}
