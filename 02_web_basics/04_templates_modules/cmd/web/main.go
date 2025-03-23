package main

import (
	"fmt"
	"net/http"

	"github.com/bciupka/go-mod/02_web_basics/04_templates_modules/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Server starting on port", port)
	http.ListenAndServe(port, nil)
}
