package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Server starting on port", port)
	http.ListenAndServe(port, nil)
}
